package item

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/util"
	"github.com/tteeoo/mudcord/item/consumable"
)

// Item represents a weapon, hat, ammo, quest item, or consumable
type Item interface {
	Inspect() []*discordgo.MessageEmbedField
	ID() string
	Desc() string
	Display() string
	Price() int
	Type() string
	CombatUsable() bool
	Use(*util.Context)
}

// Items contains a string mapped to each item
var Items map[string]Item = map[string]Item{
	consumable.Canteen.ID(): consumable.Canteen,
}
