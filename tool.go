package client

// ParameterType determines the data type of a parameter.
type ParameterType string

const (
	// ParameterTypeString indicates the parameter accepts string/text values.
	ParameterTypeString ParameterType = "string"
)

// BodyEncoding determines how the HTTP body will be encoded.
type BodyEncoding string

const (
	// BodyEncodingForm indicates the body will be encoded using URL encoding.
	BodyEncodingForm BodyEncoding = "application/x-www-form-urlencoded"

	// BodyEncodingJSON indicates the body will be encoded as JSON.
	BodyEncodingJSON BodyEncoding = "application/json"
)

type ToolList struct {
	Tools []*Tool `json:"tools"`
}

type Tool struct {
	ID          string                    `json:"id"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Parameters  []ToolParameter           `json:"parameters"`
	Webhook     *ToolWebhookConfiguration `json:"webhook,omitempty"`
	HTTP        *ActionHTTPDefinition     `json:"http,omitempty"`
	Mock        bool                      `json:"mock,omitempty"`
}

type ToolParameter struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        ParameterType     `json:"type"`
	Required    *bool             `json:"required,omitempty"`
	Options     []ParameterOption `json:"options,omitempty"`
}

type ParameterOption struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}

type ToolWebhookConfiguration struct {
	Name string `json:"name"`
}

type ActionHTTPDefinition struct {
	Method          string                    `json:"method"`
	URLTemplate     string                    `json:"url_template"`
	HeaderTemplates map[string]string         `json:"header_templates,omitempty"`
	Body            *ActionHTTPBodyDefinition `json:"body,omitempty"`
}

type ActionHTTPBodyDefinition struct {
	Encoding           BodyEncoding      `json:"encoding"`
	JSONTemplate       string            `json:"json_template,omitempty"`
	FormFieldTemplates map[string]string `json:"form_field_templates,omitempty"`
}
