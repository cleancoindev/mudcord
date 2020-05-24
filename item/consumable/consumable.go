package consumable

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/util"
	"github.com/tteeoo/mudcord/db"
	"strconv"
)

type Consumable struct {
	Price             int
	Desc, Display, ID string
	CombatUsable      bool
	consume           func(*util.Context)
}

func (item Consumable) Inspect() []*discordgo.MessageEmbedField {
	return []*discordgo.MessageEmbedField{
		{Name: "Type", Value: "Consumable", Inline: true},
		{Name: "Sell price", Value: strconv.Itoa(item.Price), Inline: true},
		{Name: "Combat usable", Value: strconv.FormatBool(item.CombatUsable), Inline: true},
	}
}

func (item Consumable) Use(ctx *util.Context) {
	item.consume(ctx)
	user, _ := db.GetUser(ctx.Message.Author.ID)
	user.RemoveItem(item.ID)
	db.SetUser(user)
}
