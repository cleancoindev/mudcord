package weapon

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
	"strconv"
)

type Weapon struct {
	weaponType, desc, display, ID            string
	Price, Damage, Accuracy, Crit, Speed, MP int
}

func (item Weapon) Inspect() []*discordgo.MessageEmbedField {
	return []*discordgo.MessageEmbedField{
		{Name: "Type", Value: "Weapon (" + item.weaponType + ")", Inline: true},
		{Name: "Sell price", Value: strconv.Itoa(item.Price), Inline: true},
		{Name: "Damage", Value: strconv.Itoa(item.Damage), Inline: true},
		{Name: "Accuracy", Value: strconv.Itoa(item.Accuracy), Inline: true},
		{Name: "Crit %", Value: strconv.Itoa(item.Crit), Inline: true},
		{Name: "Speed", Value: strconv.Itoa(item.Speed), Inline: true},
		{Name: "MP", Value: strconv.Itoa(item.MP), Inline: true},
	}
}

func (item Weapon) Use(ctx *util.Context) {
	user, _ := db.GetUser(ctx.Message.Author.ID)
	if user.AddArs(item.ID) {
		ctx.Reply("moved **" + item.display + "** from your inventory to your weapons arsenal")
		user.RemoveItem(item.ID)
		db.SetUser(user)
		return
	}

	ctx.Reply("your weapons arsenal is full")
}

func (item Weapon) Desc() string {
	return item.desc
}

func (item Weapon) Display() string {
	return item.display
}
