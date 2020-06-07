package character

import (
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
	"github.com/tteeoo/mudcord/item"
)

const HatrmHelp = "hatrm; Dequips your hat"

func Hatrm(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)

	// Send message if empty
	if user.Hat == "NoneHat" {
		ctx.Reply("you are not wearing a hat")
		return
	}

	// Change hats
	ctx.Reply("dequipped **" + item.Items[user.Hat].Display() + "**")
	user.AddItem(user.Hat, 1)
	user.Wear("NoneHat")

	db.SetUser(user)
}
