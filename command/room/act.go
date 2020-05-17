package command

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const ActHelp = "act <action#>; does a room specific action"

func Act(ctx *util.Context) {

	// return and send message if character is not started
	if !data.CheckStarted(ctx.Message.Author.ID) {
		NoneDialog(s, m)
		return
	}

	user := data.Users[ctx.Message.Author.ID]

	// Cannot use in combat
	if user.Combat {
		ctx.Reply(util.NoneCombat)
		return
	}

	// Get the players current room
	room := Rooms[user.Room]

	// Get action number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		ctx.Reply("that action does not exist")
		return
	}
	num--

	// return if action number does not exist
	if num <= -1 || len(room.Actions) <= num {
		ctx.Reply("that action does not exist")
		return
	}
	room.Actions[num].Fn(s, m)
}
