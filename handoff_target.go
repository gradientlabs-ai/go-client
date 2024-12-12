package client

type HandOffTarget struct {
	// ID is your identifier of choice for this hand-off target. Can be anything consisting
	// of letters, numbers, or any of the following characters: `_` `-` `+` `=`.
	ID string `json:"id"`

	// Name is the hand-off target’s name.
	Name string `json:"name"`
}
