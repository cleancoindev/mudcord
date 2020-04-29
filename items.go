package main

// Item represents a generic item
type Item struct {
	Type string
}

// Items contains all items (to be stored in json nicer as strings)
var Items map[string]Item = map[string]Item{
	"HatNone": HatNone,
}

// HatNone is used when a character has no hat
var HatNone Item = Item{Type: "hat"}
