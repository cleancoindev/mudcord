package main

// Item represents a generic item
type Item struct {
	Type string
}

var (
	// Items contains all items (to be stored in json nicer as strings)
	Items map[string]Item = map[string]Item{
		"HatNone": HatNone,
	}

	// HatNone is used when a character has no hat
	HatNone Item = Item{Type: "hat"}
)
