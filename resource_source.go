package client

import "time"

// ResourceSource represents a resource source in the system.
type ResourceSource struct {
	// ID is a generated ID for the resource source.
	ID string `json:"id"`

	// DisplayName free-text field to describe the source, eg. "Intercom User Attributes". Enforced unique per
	// company.
	DisplayName string `json:"display_name"`

	// Description optional free-text field to describe the source in more detail. Not used by the agent at all, only
	// humans.
	Description string `json:"description"`

	// SourceType describes how the data is fetched from the source.
	SourceType SourceType `json:"source_type"`

	// HTTPConfig can be set up by customers.
	HTTPConfig *ResourceHTTPDefinition `json:"http_config,omitempty"`

	// WebhookConfig can be set up by customers.
	WebhookConfig *ResourceWebhookDefinition `json:"webhook_config,omitempty"`

	// AttributeDescriptions optional raw attribute-level descriptions, used when generating the schema and as
	// additional information for the agent.
	// - key: a JSONPath eg. `$.name` or `$.items[*].name`
	// - value: a description of the attribute
	AttributeDescriptions map[string]string `json:"attribute_descriptions,omitempty"`

	// Schema is the schema of the resource source, inferred from the source payloads. It is updated
	// asynchronously as data is fetched from the source. Nil if the schema has not been inferred yet. Includes
	// attribute descriptions if present.
	Schema *Schema `json:"schema,omitempty"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type SourceType string

const (
	SourceTypeHTTP     SourceType = "http"
	SourceTypeInternal SourceType = "internal"
	SourceTypeWebhook  SourceType = "webhook"
)

// ResourceHTTPBodyDefinition determines how the HTTP request body is constructed.
type ResourceHTTPBodyDefinition struct {
	// Encoding determines how the HTTP request body will be encoded.
	Encoding string `json:"encoding"`

	// JSONTemplate contains a template that will be used to generate JSON for the
	// HTTP request body. Only used when Encoding is "application/json".
	JSONTemplate string `json:"json_template,omitempty"`

	// FormFieldTemplates contains templates for the values that will be form-encoded
	// and used as the HTTP body. Only used when Encoding is "application/x-www-form-urlencoded".
	FormFieldTemplates map[string]string `json:"form_field_templates,omitempty"`
}

// ResourceHTTPDefinition contains configuration for HTTP actions.
type ResourceHTTPDefinition struct {
	// Method is the HTTP request method that will be used.
	Method string `json:"method"`

	// URLTemplate contains a template used to construct the request URL.
	URLTemplate string `json:"url_template"`

	// Body determines how the HTTP request body is constructed.
	Body *HTTPBodyDefinition `json:"body,omitempty"`

	// HeaderTemplates contains templates for the values that will be used as request headers.
	HeaderTemplates map[string]string `json:"header_templates,omitempty"`
}

// ResourceWebhookDefinition contains configuration for webhook actions.
type ResourceWebhookDefinition struct {
	// Name will be included in the `data.action` field of the webhook payload.
	Name string `json:"name"`
}
