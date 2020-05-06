package main

import (
	"github.com/bwmarrin/discordgo"
)

// Item represents a generic item
type Item struct {
	Type    string
	Display string
	Desc    string
	Usable  bool
	Use     func(item Item, s *discordgo.Session, m *discordgo.MessageCreate)
}

var (
	// Items contains all items (to be stored in json nicer as strings)
	Items map[string]Item = map[string]Item{
		"HatNone":     HatNone,
		"ItemCanteen": ItemCanteen,
		"ItemShard":   ItemShard,
	}

	// HatNone is used when a character has no hat
	HatNone Item = Item{Type: "Hat", Display: "", Desc: "", Usable: false, Use: NoUse}

	// ItemCanteen is a canteen
	ItemCanteen Item = Item{Type: "Item", Display: "Canteen", Desc: "Glug glug glug", Usable: true, Use: NoUse}

	// ItemShard is a simplex shard
	ItemShard Item = Item{Type: "Item", Display: "Simplex Shard", Desc: "Shiny", Usable: false, Use: NoUse}
)
