package command

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const GoHelp = "go <room#>; goes to a specific room"

func Go(ctx *util.Context) {

	// return and send message if character is not started
	if !data.CheckStarted(m.Author.ID) {
		NoneDialog(s, m)
		return
	}

	user := data.Users[ctx.Message.Author.ID]

	// Cannot use in combat
	if user.Combat {
		NoneCombat(s, m)
		return
	}

	// Get the players current room
	room := Rooms[user.Room]

	// Get room number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		ctx.Reply("that room does not exist")
		return
	}
	num--

	// return if room number does not exist and change player room
	if num <= -1 || len(room.Rooms) <= num {
		ctx.Reply("that room does not exist")
		return
	}

	ctx.Reply(Rooms[room.Rooms[num]].Into)
	user.Room = room.Rooms[num]
}
