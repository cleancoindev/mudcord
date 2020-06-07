package character

import (
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/util"
)

const HatrmHelp = "hatrm; Dequips your hat"

func Hatrm(ctx *util.Context) {

	user, err := db.GetUser(ctx.Message.Author.ID)
	if util.CheckDB(err, ctx) {
		return
	}

	// Send message if empty
	if user.Hat == "None" {
		ctx.Reply("you are not wearing a hat")
		return
	}

	// Change hats
	ctx.Reply("dequipped **" + item.Items[user.Hat].Display() + "**")
	user.AddItem(user.Hat, 1)
	user.Wear("None")

	err = db.SetUser(user)
	util.CheckDB(err, ctx)
}
