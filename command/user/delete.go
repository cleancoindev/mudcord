package command

import (
	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const DeleteHelp = "delete; deletes your character data"

func Delete(ctx *util.Context) {

	// Ensure command author has started their journey
	if !data.CheckStarted(ctx.Message.Author.ID) {
		ctx.Reply("you do not have a character to delete")
		return
	}

	// Delete the authors info from the Users map
	delete(data.Users, ctx.Message.Author.ID)
	ctx.Reply("successfully deleted your character, run `start` to start a new one")
}
