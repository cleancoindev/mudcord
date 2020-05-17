package command

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const TalkHelp = "talk <npc#>; talks to an npc"

func Talk(ctx *util.Context) {

	// return and send message if character is not started
	if !data.CheckStarted(ctx.Message.Author.ID) {
		ctx.Reply(util.NoneDialog)
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

	// Get npc number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		ctx.Reply("that npc does not exist")
		return
	}
	num--

	// return if npc number does not exist
	if num <= -1 || len(room.NPCs) <= num {
		ctx.Reply("that npc does not exist")
		return
	}

	ctx.Reply("**" + room.NPCs[num].Name + ":** " + room.NPCs[num].Speak(room.NPCs[num]))
}
