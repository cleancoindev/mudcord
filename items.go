package main

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
)

// Item represents a generic item
type Item struct {
	Price                  int
	Desc, Display, useName string
	hat, weapon            bool
	Usable, CombatUsable   bool
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
	ItemCanteen Item = Item{
		Display:      "Canteen",
		Desc:         "A shiny, refillable container; heals up to two HP",
		Usable:       true,
		CombatUsable: true,
		useName:      "Canteen",
	}

	// ItemEmptyCanteen is an empty canteen
	ItemEmptyCanteen Item = Item{
		Display: "Empty Canteen",
		Desc:    "A shiny, refillable container",
	}

	// Hats

	// HatNone is used when a character has no hat
	HatNone Item = Item{
		hat: true,
	}
)

// Uses

// Use runs code to use a specific item
func (item Item) Use(num int, s *discordgo.Session, m *discordgo.MessageCreate) {

	if !item.Usable {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that item cannot be used")
		return
	}

	user := Users[m.Author.ID]

	if !item.CombatUsable && user.Combat {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that item cannot be used in combat")
		return
	}

	switch item.useName {
	case "Canteen":
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

}

// Type returns the type of an item as a string
func (item Item) Type() string {
	if item.hat {
		return "Hat"
	}
	if item.weapon {
		return "Weapon"
	}
	return "Item"
}
