package character

import (
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
)

const DeleteHelp = "delete; deletes your character data"

func Delete(ctx *util.Context) {

	db.DeleteUser(ctx.Message.Author.ID)
	ctx.Reply("successfully deleted your character, run `start` to start a new one")
}
