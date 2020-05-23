package option

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const TalkHelp = "talk <npc#>; talks to an npc"

func Talk(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)

	// Get the players current room
	currentRoom := room.Rooms[user.Room]

	// Get npc number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		ctx.Reply("that npc does not exist")
		return
	}
	num--

	// return if npc number does not exist
	if num <= -1 || len(currentRoom.NPCs) <= num {
		ctx.Reply("that npc does not exist")
		return
	}

	ctx.Reply("**" + currentRoom.NPCs[num].Name + ":** " + currentRoom.NPCs[num].Speak(currentRoom.NPCs[num]))
}
