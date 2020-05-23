package character

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/util"
)

const UseHelp = "use <item#>; uses an item from your inventory"

func Use(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)

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

	currentItem := item.Items[user.Inv[num].ID]

	currentItem.Use(ctx)
}
