package client

import "encoding/json"

// Schema of data related to a resource.
type Schema struct {
	// Raw is the raw schema data.
	// - inferred from the data payloads
	// - updated asynchronously as data is fetched
	// - can be unmarshalled into a jsonschema.Schema (use Parse method)
	Raw json.RawMessage `json:"raw"`

	// Attributes is a list of attributes derived from the schema. Only leaf properties are included in this list,
	// container objects and arrays themselves are not included as separate attributes.
	//
	// For example, given a schema with:
	//   {
	//     "properties": {
	//       "name": { "type": "string" },
	//       "address": {
	//         "type": "object",
	//         "properties": {
	//           "street": { "type": "string" },
	//           "city": { "type": "string" }
	//         }
	//       },
	//       "skills": {
	//         "type": "array",
	//         "items": {
	//           "type": "object",
	//           "properties": {
	//             "name": { "type": "string" },
	//             "level": { "type": "string" }
	//           }
	//         }
	//       }
	//     }
	//   }
	//
	// The resulting attributes would be:
	//   - $.name (string)
	//   - $.address.street (string)
	//   - $.address.city (string)
	//   - $.skills[*].name (string)
	//   - $.skills[*].level (string)
	//
	// Note that $.address and $.skills are not included as separate attributes.
	// Attributes are sorted alphabetically by JSONPath.
	Attributes []Attribute `json:"attributes"`
}
