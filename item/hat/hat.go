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
		{Name: "Type", Value: "Hat"},
		{Name: "Sell price", Value: strconv.Itoa(item.Price), Inline: true},
		{Name: "Defense", Value: strconv.Itoa(item.Def), Inline: true},
		{Name: "Resistance", Value: strconv.Itoa(item.Res), Inline: true},
	}
}

// hats is need to get old hat names without import cycles
var hats = map[string]Hat{
	Bucket.ID: Bucket,
}

func (item Hat) Use(ctx *util.Context) {
	user, err := db.GetUser(ctx.Message.Author.ID)
	if util.CheckDB(err, ctx) {
		return
	}

	oldHat := user.Wear(item.ID)
	user.RemoveItem(item.ID)

	if oldHat != "None" {
		ctx.Reply("equipped **" + item.display + "** and unequipped **" + hats[oldHat].display + "**")
		user.AddItem(oldHat, 1)
		err = db.SetUser(user)
		util.CheckDB(err, ctx)
		return
	}

	ctx.Reply("equipped **" + item.display + "**")
	err = db.SetUser(user)
	util.CheckDB(err, ctx)
}

func (item Hat) Desc() string {
	return item.desc
}

func (item Hat) Display() string {
	return item.display
}
