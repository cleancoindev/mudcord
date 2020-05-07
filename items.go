package main

import (
	"github.com/bwmarrin/discordgo"
)

// Item represents a generic item
type Item struct {
	Value   int
	Type    string
	Display string
	Desc    string
	Usable  bool
	Use     func(num int, s *discordgo.Session, m *discordgo.MessageCreate)
}

// ItemQuan represents an item string with a quantity value
type ItemQuan struct {
	Item string
	Quan int
}

var (
	// Items contains all items (to be stored in json nicer as strings)
	Items map[string]Item = map[string]Item{
		"HatNone":     HatNone,
		"ItemCanteen": ItemCanteen,
		"ItemShard":   ItemShard,
	}

	// Items

	// ItemCanteen is a canteen
	ItemCanteen Item = Item{Type: "Item", Display: "Canteen", Desc: "A shiny container; heals one HP", Usable: true, Use: UseCanteen}

	// ItemShard is a simplex shard
	ItemShard Item = Item{Type: "Item", Display: "Simplex Shard", Desc: "Shiny", Usable: false, Use: UseNone}

	// Hats

	// HatNone is used when a character has no hat
	HatNone Item = Item{Type: "Hat", Display: "", Desc: "", Usable: false, Use: UseNone}
)

// Uses

// UseNone is for items that cannot be used
func UseNone(_ int, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that item cannot be used")
}

// UseCanteen uses the canteen item
func UseCanteen(num int, s *discordgo.Session, m *discordgo.MessageCreate) {

	user := Users[m.Author.ID]
	user.RemoveItem(num)

}
