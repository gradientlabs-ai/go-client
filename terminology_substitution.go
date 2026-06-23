package client

import "time"

// TerminologySubstitution configures a word or phrase that the AI agent should
// avoid, replacing it with an approved alternative.
type TerminologySubstitution struct {
	// ID uniquely identifies this terminology substitution.
	ID string `json:"id"`

	// Blocked is the word or phrase the AI agent must not use.
	Blocked string `json:"blocked"`

	// BlockedDescription explains why the blocked term should not be used.
	BlockedDescription string `json:"blocked_description"`

	// Replacement is the approved alternative the AI agent should use instead.
	Replacement string `json:"replacement"`

	// ResourceTypeID optionally scopes the substitution to a specific resource type.
	ResourceTypeID string `json:"resource_type_id,omitempty"`

	// ResourceAttributeJSONPath optionally scopes the substitution to a specific
	// attribute within the resource type.
	ResourceAttributeJSONPath string `json:"resource_attribute_json_path,omitempty"`

	// ResourceValueToMatch optionally further scopes the substitution to cases where
	// the resource attribute equals this value.
	ResourceValueToMatch string `json:"resource_value_to_match,omitempty"`

	// Created is the time at which this substitution was created.
	Created time.Time `json:"created"`

	// Updated is the time at which this substitution was last updated.
	Updated time.Time `json:"updated"`
}
