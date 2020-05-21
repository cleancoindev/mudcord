package thing

import (
	"strconv"
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/util"
	"github.com/tteeoo/mudcord/thing/consumable"
)

// Item represents a weapon, hat, ammo, quest item, or consumable
type Item interface {
	Use(*util.Context)
	Type() string
	Inspect() []*discordgo.MessageEmbedField
}

// ItemQuan represents an item id with a quantity value
type ItemQuan struct {
	ID string
	Quan int
}

// Items contains a string mapped to each item
var Items map[string]Item = map[string]Item{
	consumable.Canteen.ID: consumable.Canteen,
}

// // Use runs code to use a specific item
// func (item Item) Use(num int, s *discordgo.Session, m *discordgo.MessageCreate) {

// 	user := Users[m.Author.ID]

// 	if item.weapon {
// 		user.AddArs(num, s, m)
// 		return
// 	}

// 	if !item.Usable {
// 		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that item cannot be used")
// 		return
// 	}

// 	if !item.CombatUsable && user.Combat {
// 		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that item cannot be used in combat")
// 		return
// 	}

// 	switch item.ID {
// 	case "ItemCanteen":

// 		healed := user.Heal(2)
// 		if healed == 0 {
// 			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you are already at full health")
// 			return
// 		}

// 		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you chug down the water inside, healing "+strconv.Itoa(healed)+" health")
// 		user.RemoveItem(num)
// 		user.AddItem("ItemEmptyCanteen", 1)
// 	}

// }
