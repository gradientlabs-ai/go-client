package client

// Attribute describes an attribute available when accessing a resource of a given type.
type Attribute struct {
	// Path expression to read this attribute from a resource payload.
	Path string

	// Type describes the data type of the attribute.
	Type AttributeType

	// Cardinality determines how many values would be returned when applying the
	// given JSONPath. Only attributes with a cardinality of "one" can be used as
	// tool parameter values.
	Cardinality AttributeCardinality

	// Description is an optional human-readable description of the attribute.
	Description string

	// IsTopLevel is true if the attribute is a top-level field (direct child of the root).
	IsTopLevel bool

	// Name is the name of the field, e.g.:
	//   - "id" for `$.id`
	//   - "quantity" for `$.items[*].quantity`
	Name string
}

// AttributeCardinality determines how many values would be returned when
// applying an attribute's JSONPath to a resource payload.
type AttributeCardinality string

const (
	// AttributeCardinalityOne means the attribute represents one value (e.g.
	// `$.id` or `$.name`).
	AttributeCardinalityOne AttributeCardinality = "one"

	// AttributeCardinalityMany means the attribute represents many values (e.g.
	// `$.items[*].quantity`).
	AttributeCardinalityMany AttributeCardinality = "many"
)

// AttributeType describes the data type of the attribute.
type AttributeType string

const (
	// AttributeTypeString means the attribute is a string.
	AttributeTypeString AttributeType = "string"

	// AttributeTypeDate means the attribute is a date in the string form:
	// YYYY-MM-DD.
	AttributeTypeDate AttributeType = "date"

	// AttributeTypeTimestamp means the attribute is a timestamp in the RFC3339 form.
	AttributeTypeTimestamp AttributeType = "timestamp"

	// AttributeTypeBoolean means the attribute is a boolean value.
	AttributeTypeBoolean AttributeType = "boolean"

	// AttributeTypeNumber means the attribute is a numeric value. JSON doesn't
	// differentiate between integers and floating point numbers, so it's up to
	// the user to interpret this.
	AttributeTypeNumber AttributeType = "number"

	// AttributeTypeArray means the attribute is an array of primitive values.
	//
	// Note: an array of primitive values will still have a cardinality of "one"
	// because there is one value available (the array itself).
	AttributeTypeArray AttributeType = "array"

	// AttributeTypeComplex means the attribute could be many types, which we don't
	// support yet.
	AttributeTypeComplex AttributeType = "complex"
)
