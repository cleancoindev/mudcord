package main

// Item represents a generic item
type Item struct {
	Hat bool
}

// HatNone is used when a character has no hat
var HatNone Item = Item{Hat: true}
