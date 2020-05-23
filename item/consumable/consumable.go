package consumable

import (
	"strconv"
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/util"
)

type Consumable struct {
	price             int
	desc, display, id string
	combatUsable      bool
	consume               func(*util.Context)
}

func (item Consumable) Inspect() []*discordgo.MessageEmbedField {
    return []*discordgo.MessageEmbedField{
		{Name: "Type", Value: "Consumable", Inline: true},
		{Name: "Sell price", Value: strconv.Itoa(item.Price()), Inline: true},
		{Name: "Combat usable", Value: strconv.FormatBool(item.CombatUsable()), Inline: true},
	}
}

func (item Consumable) Type() string {
	return "Consumable"
}

func (item Consumable) Desc() string {
	return item.desc
}

func (item Consumable) Display() string {
	return item.display
}

func (item Consumable) ID() string {
	return item.id
}

func (item Consumable) CombatUsable() bool {
	return item.combatUsable
}

func (item Consumable) Price() int {
	return item.price
}

func (item Consumable) Use(ctx *util.Context) {
	item.consume(ctx)
}
