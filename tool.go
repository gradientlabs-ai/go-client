package client

import "time"

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
	HTTP        *HTTPDefinition           `json:"http,omitempty"`
	Async       *AsyncDefinition          `json:"async,omitempty"`
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

type HTTPDefinition struct {
	Method          string              `json:"method"`
	URLTemplate     string              `json:"url_template"`
	HeaderTemplates map[string]string   `json:"header_templates,omitempty"`
	Body            *HTTPBodyDefinition `json:"body,omitempty"`
}

type HTTPBodyDefinition struct {
	Encoding           BodyEncoding      `json:"encoding"`
	JSONTemplate       string            `json:"json_template,omitempty"`
	FormFieldTemplates map[string]string `json:"form_field_templates,omitempty"`
}

type AsyncDefinition struct {
	// StartExecutionTool is the tool that will be executed to start the async operation. It should return a result
	// fairly quickly that optionally contains information about the work that has been kicked off. It is executed like
	// a normal HTTP/Webhook tool.
	StartExecutionTool ChildTool `json:"start_execution_tool"`

	// Timeout is the maximum time the async operation is allowed to run. If it exceeds this time, we will stop waiting
	// for the response and the agent will be notified that the async operation has timed out.
	Timeout time.Duration `json:"timeout"`
}

// ChildTool represents a tool that is part of a parent tool's configuration. It does not have its own ID, company ID,
// version, dates, authors, etc. as these are all inherited from the parent tool. In its current form it effectively
// maps to an action definition.
//
// Note that the child tool does not currently support parameters, these are decided by the parent tool and passed to
// the child tool when it is executed.
type ChildTool struct {
	Webhook *ToolWebhookConfiguration `json:"webhook,omitempty"`
	HTTP    *HTTPDefinition           `json:"http,omitempty"`
}
