package item

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/item/consumable"
	"github.com/tteeoo/mudcord/item/weapon"
	"github.com/tteeoo/mudcord/util"
)

// Item represents a weapon, hat, ammo, quest item, or consumable
type Item interface {
	Inspect() []*discordgo.MessageEmbedField
	Use(*util.Context)
}

// Items contains a string mapped to each item
var Items = map[string]Item{
	consumable.Canteen.ID: consumable.Canteen,
	weapon.WoodSword.ID:   weapon.WoodSword,
}
