package character

import (
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
)

const DeleteHelp = "delete; deletes your character data"

func Delete(ctx *util.Context) {

	err := db.DeleteUser(ctx.Message.Author.ID)
	if util.CheckDB(err, ctx) {
		return
	}
	ctx.Reply("successfully deleted your character, run `start` to start a new one")
}
