package client

import "time"

type ResourceType struct {
	// ID is a generated ID for the resource type.
	ID string `json:"id"`

	// DisplayName is a freetext human-readable name for the resource type.
	DisplayName string `json:"display_name"`

	// Description is an optional freetext human-readable description of the resource type.
	Description string `json:"description"`

	// Scope determines when in the conversation the resource is fetched and used.
	Scope Scope `json:"scope"`

	// RefreshStrategy determines how often the resource is re-fetched.
	RefreshStrategy RefreshStrategy `json:"refresh_strategy"`

	// SourceConfig defines how the resource is fetched. If not set, the type can still be used in testing environments
	// with mocked data.
	SourceConfig *SourceConfig `json:"source_config,omitempty"`

	// Schema is the JSON schema for the resource type, optionally including attribute descriptions from the source if
	// IncludeDescriptions is true on read.
	Schema *Schema `json:"schema,omitempty"`

	// IsEnabled indicates whether the resource type has been configured as enabled or not.
	IsEnabled bool `json:"is_enabled"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type SourceConfig struct {
	// SourceID identifies the source which should be used to fetch the resource data.
	SourceID string `json:"source_id"`

	// Attributes defines the *top-level* fields names to be used from the source data. Eg. for a payload like this:
	//
	// ```
	// {
	//   "id": 123,
	//   "name": "foo",
	//   "items": [
	//     {
	//       "id": 1,
	//       "name": "bar"
	//     },
	//     {
	//       "id": 2,
	//       "name": "baz"
	//     }
	//   ],
	//	"nested": {
	//     "id": 456,
	//     "name": "qux"
	//   }
	// }
	// ```
	//
	// The following attributes would be valid:
	//   - "id"
	//   - "name"
	//   - "items"
	//   - "nested"
	// The following attributes would be invalid:
	//   - "items[*].id"
	//   - "nested.id"
	Attributes []string `json:"attributes"`

	// Cache determines how long we'll consider data from the source be valid before refreshing it. It's either a
	// duration string (e.g. "5m") or CacheNever if the resource is refreshed on every turn. The default is 1 minute.
	Cache string `json:"cache"`
}

// CacheNever means the pull resource will be refreshed on every turn of the
// conversation.
const CacheNever = "never"

type Scope string

const (
	// ScopeGlobal means the resource is available throughout the conversation and in all procedures.
	ScopeGlobal Scope = "global"

	// ScopeLocal means the resource is available only in procedures that explicitly use it, only fetched when it's used.
	ScopeLocal Scope = "local"
)

type RefreshStrategy string

const (
	// RefreshStrategyDynamic means the resource value can change, so this is re-fetched throughout the conversation.
	RefreshStrategyDynamic RefreshStrategy = "dynamic"

	// RefreshStrategyStatic means the resource is fetched once at the start of the conversation (global) or when it's first used in a procedure (local).
	RefreshStrategyStatic RefreshStrategy = "static"
)
