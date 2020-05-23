package option

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const GoHelp = "go <room#>; goes to a specific room"

func Go(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)

	// Get the players current room
	currentRoom := room.Rooms[user.Room]

	// Get room number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		ctx.Reply("that room does not exist")
		return
	}
	num--

	// return if room number does not exist and change player room
	if num <= -1 || len(currentRoom.Rooms) <= num {
		ctx.Reply("that room does not exist")
		return
	}

	ctx.Reply(room.Rooms[currentRoom.Rooms[num]].Into)
	user.Room = currentRoom.Rooms[num]
	db.SetUser(user)
}
