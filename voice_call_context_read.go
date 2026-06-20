package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ReadVoiceCallContextParams are the parameters to Client.ReadLatestVoiceCallContext.
type ReadVoiceCallContextParams struct {
	// LookbackSeconds sets the time window (in seconds) within which to search for
	// recent call events. Defaults to 60 if not set; minimum 5.
	LookbackSeconds int

	// IncludeLargeFields includes the full transcript and complete summary in the
	// response when true. Defaults to false.
	IncludeLargeFields bool
}

// ReadLatestVoiceCallContext returns the most recent voice call context for the
// given phone number.
//
// Note: requires a Public API key.
func (c *Client) ReadLatestVoiceCallContext(ctx context.Context, phoneNumber string, p *ReadVoiceCallContextParams) (*VoiceCallContext, error) {
	path := fmt.Sprintf("voice/latest-call-context/%s", phoneNumber)
	if p != nil {
		sep := "?"
		if p.LookbackSeconds > 0 {
			path += sep + "lookback_seconds=" + strconv.Itoa(p.LookbackSeconds)
			sep = "&"
		}
		if p.IncludeLargeFields {
			path += sep + "include_large_fields=true"
		}
	}

	rsp, err := c.makeRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var result VoiceCallContext
	if err := json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
