package client

import (
	"encoding/json"
	"time"
)

// BackOfficeTaskStatus describes the current state of a back-office task.
type BackOfficeTaskStatus string

// BackOfficeTask represents a back-office task managed by the AI agent.
type BackOfficeTask struct {
	ID             string                             `json:"id"`
	AgentID        string                             `json:"agent_id"`
	Status         BackOfficeTaskStatus               `json:"status,omitempty"`
	Input          json.RawMessage                    `json:"input"`
	Metadata       map[string]string                  `json:"metadata,omitempty"`
	Attachments    []BackOfficeTaskResponseAttachment `json:"attachments,omitempty"`
	Created        time.Time                          `json:"created"`
	Updated        time.Time                          `json:"updated,omitempty"`
	Completed      time.Time                          `json:"completed,omitempty"`
	Failed         time.Time                          `json:"failed,omitempty"`
	FailureReasons []string                           `json:"failure_reasons,omitempty"`
	HandedOff      time.Time                          `json:"handed_off,omitempty"`
	HandOffReason  string                             `json:"hand_off_reason,omitempty"`
	Result         *BackOfficeTaskResult              `json:"result,omitempty"`
}

// BackOfficeTaskResponseAttachment is an attachment returned in a BackOfficeTask response.
type BackOfficeTaskResponseAttachment struct {
	IdempotencyKey string `json:"idempotency_key"`
	FileName       string `json:"file_name"`
	ExternalURL    string `json:"external_url,omitempty"`
	RawContents    []byte `json:"raw_contents,omitempty"`
}

// BackOfficeTaskResult holds the output of a completed back-office task.
type BackOfficeTaskResult struct {
	ResultType string           `json:"result_type"`
	Custom     *json.RawMessage `json:"custom,omitempty"`
}
