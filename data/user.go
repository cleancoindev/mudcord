package data

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/thing"
)

// User represents a character
type User struct {
	Level, Gold, XP int
	HP, MP          [2]int
	Combat          bool
	Room, Hat       string
	Inv             []*ItemQuan
	Arsenal         []string
}

// RemoveItem will either remove an item from a users inventory or decrement the quantity
func (user *User) RemoveItem(index int) {
	if user.Inv[index].Quan > 1 {
		user.Inv[index].Quan--
		return
	}

	user.Inv = append(user.Inv[:index], user.Inv[index+1:]...)
}

// AddItem will either add an item to a users inventory or increment the quantity
func (user *User) AddItem(item string, quan int) {
	for _, val := range user.Inv {
		if val.Item == item {
			val.Quan++
			return
		}
	}

	user.Inv = append(user.Inv, &ItemQuan{Item: item, Quan: quan})
}

// AddArs will attempt to add an item to a users arsenal
func (user *User) AddArs(index int, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(user.Arsenal) >= 3 {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" your weapon arsenal is full")
		return
	}

	item := Items[user.Inv[index].Item]

	user.Arsenal = append(user.Arsenal, item.ID)
	user.RemoveItem(index)
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" moved **"+item.Display+"** to your weapon arsenal")
}

// RemoveArs will attempt to remove an item to a users arsenal
func (user *User) RemoveArs(index int, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(user.Arsenal) <= index {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that weapon does not exist")
		return
	}

	item := Items[user.Arsenal[index]]

	user.Arsenal = append(user.Arsenal[:index], user.Arsenal[index+1:]...)
	user.AddItem(item.ID, 1)
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" moved **"+item.Display+"** out of your weapon arsenal")
}

// InvCount gets the total number of items in a users inventory
func (user *User) InvCount() int {
	var count int
	for _, val := range user.Inv {
		count += val.Quan
	}

	return count
}

// Heal attempts to heal the user, taking into account max hp and returns the amount healed
func (user *User) Heal(amount int) int {
	if user.HP[0] >= user.HP[1] {
		return 0
	}

	if user.HP[0]+amount > user.HP[1] {
		user.HP[0] = user.HP[1]
		return user.HP[1] - user.HP[0]
	}

	user.HP[0] += amount
	return amount
}
