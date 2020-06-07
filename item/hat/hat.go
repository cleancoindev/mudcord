package hat

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
	"strconv"
)

type Hat struct {
	desc, display, ID string
	Price, Def, Res   int
}

func (item Hat) Inspect() []*discordgo.MessageEmbedField {
	return []*discordgo.MessageEmbedField{
		{Name: "Sell price", Value: strconv.Itoa(item.Price), Inline: true},
		{Name: "Defense", Value: strconv.Itoa(item.Def), Inline: true},
		{Name: "Resistance", Value: strconv.Itoa(item.Res), Inline: true},
	}
}

func (item Hat) Use(ctx *util.Context) {
	user, _ := db.GetUser(ctx.Message.Author.ID)
	oldHat := user.Wear(item.ID)
	user.RemoveItem(item.ID)

	if oldHat != "NoneHat" {
		ctx.Reply("equipped **" + item.display + "** and unequipped **" + oldHat + "**")
		user.AddItem(oldHat, 1)
		db.SetUser(user)
		return
	}

	ctx.Reply("equipped **" + item.display + "**")
	db.SetUser(user)
}

func (item Hat) Desc() string {
	return item.desc
}

func (item Hat) Display() string {
	return item.display
}
