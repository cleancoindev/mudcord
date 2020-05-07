package main

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
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
		"HatNone":          HatNone,
		"ItemCanteen":      ItemCanteen,
		"ItemEmptyCanteen": ItemEmptyCanteen,
	}

	// Items

	// ItemCanteen is a canteen
	ItemCanteen Item = Item{Type: "Item", Display: "Canteen", Desc: "A shiny, refillable container; heals up to two HP", Usable: true, Use: UseCanteen}

	// ItemEmptyCanteen is an empty canteen
	ItemEmptyCanteen Item = Item{Type: "Item", Display: "Empty Canteen", Desc: "A shiny, refillable container", Usable: false, Use: UseNone}

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

	healed := user.Heal(2)
	if healed == 0 {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you are already at full health")
		return
	}

	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you chug down the water inside, healing "+strconv.Itoa(healed)+" health")
	user.RemoveItem(num)
	user.AddItem("ItemEmptyCanteen", 1)
}
