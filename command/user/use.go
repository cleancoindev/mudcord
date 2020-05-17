package command

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const UseHelp = "use <item#>; uses an item from your inventory"

func Use(ctx *util.Context) {

	// return and send message if character is not started
	if !data.CheckStarted(ctx.Message.Author.ID) {
		ctx.Reply(util.NoneDialog)
		return
	}

	user := data.Users[ctx.Message.Author.ID]

	// Get item number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		ctx.Reply("that item does not exist")
		return
	}
	num--

	// return if item number does not exist
	if num <= -1 || len(user.Inv) <= num {
		ctx.Reply("that item does not exist")
		return
	}

	item := Items[user.Inv[num].Item]

	item.Use(num, ctx)
}
