package character

import (
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
)

const StartHelp = "start; creates a new character"

func Start(ctx *util.Context) {

	// Create a new character in the Users map
	_, err := db.NewUser(ctx.Message.Author.ID)
	if util.CheckDB(err, ctx) {
		return
	}
	ctx.Reply("you hear the captain say 'land ho!'. You've arrived at the forgotten island of Alkos, ready to start a new journey")
}
