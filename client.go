// Package client implements Go bindings for the [Gradient Labs API].
//
// [Gradient Labs API]: https://api-docs.gradient-labs.ai
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

const defaultURL = "https://api.gradient-labs.ai"

var userAgent = fmt.Sprintf("Gradient-Labs-Go (%s/%s)", version, runtime.Version())

// Client provides access to the Gradient Labs API. Use NewClient to create one.
type Client struct {
	url, apiKey     string
	httpClient      *http.Client
	webhookVerifier *WebhookVerifier
}

// NewClient creates a client with the given options.
//
// Note: the WithAPIKey option is required, as is WithWebhookSigningKey if you
// intend to receive webhooks.
func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		url:        defaultURL,
		httpClient: &http.Client{},
		webhookVerifier: &WebhookVerifier{
			leeway: defaultLeeway,
		},
	}

	for _, opt := range opts {
		switch t := opt.(type) {
		case urlOption:
			c.url = t.url
		case apiKeyOption:
			c.apiKey = t.apiKey
		case transportOption:
			c.httpClient.Transport = t.transport
		case webhookSigningKeyOption:
			c.webhookVerifier.secret = t.signingKey
		case webhookLeewayOption:
			c.webhookVerifier.leeway = t.leeway
		}
	}

	return c, nil
}

// WithURL overrides the default base URL.
func WithURL(url string) Option {
	return urlOption{url}
}

// WithTransport customises the HTTP client's Transport, which is useful for
// observability (e.g. capturing trace spaces, metrics) and in unit tests.
func WithTransport(rt http.RoundTripper) Option {
	return transportOption{rt}
}

// WithAPIKey sets the client's API key.
func WithAPIKey(key string) Option {
	return apiKeyOption{key}
}

// WithWebhookSigningKey sets the client's webhook signing key. It will be used
// to verify webhook authenticity.
func WithWebhookSigningKey(key string) Option {
	return webhookSigningKeyOption{[]byte(key)}
}

// WithWebhookLeeway determines the maximum age of webhook that will be accepted
// by Client.VerifyWebhookRequest.
func WithWebhookLeeway(leeway time.Duration) Option {
	return webhookLeewayOption{leeway}
}

type urlOption struct{ url string }
type transportOption struct{ transport http.RoundTripper }
type apiKeyOption struct{ apiKey string }
type webhookSigningKeyOption struct{ signingKey []byte }
type webhookLeewayOption struct{ leeway time.Duration }

func (urlOption) isClientOption()               {}
func (transportOption) isClientOption()         {}
func (apiKeyOption) isClientOption()            {}
func (webhookSigningKeyOption) isClientOption() {}
func (webhookLeewayOption) isClientOption()     {}

// Option customises the Client.
type Option interface{ isClientOption() }

func (c *Client) makeRequest(ctx context.Context, method string, path string, body any) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.url, path)

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", userAgent)

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return c.httpClient.Do(req)
}

func (c *Client) handleResponse(response *http.Response, value interface{}) error {
	defer response.Body.Close()

	if err := responseError(response); err != nil {
		return err
	}

	// If result is nil, we don't need to decode the response body.
	if value == nil {
		return nil
	}

	return json.NewDecoder(response.Body).Decode(value)
}
