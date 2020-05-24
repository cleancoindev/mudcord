package option

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const ActHelp = "act <action#>; does a room specific action"

func Act(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)

	// Get the players current room
	currentRoom := room.Rooms[user.Room]

	// Get action number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		ctx.Reply("that action does not exist")
		return
	}
	num--

	// return if action number does not exist
	if num <= -1 || len(currentRoom.Actions) <= num {
		ctx.Reply("that action does not exist")
		return
	}
	currentRoom.Actions[num].Fn(ctx)
}
