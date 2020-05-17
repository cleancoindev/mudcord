package command

import (
	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const StartHelp = "start; creates a new character"

func Start(ctx *util.Context) {

	// Ensure command author has not started their journey
	if data.CheckStarted(ctx.Message.Author.ID) {
		ctx.Reply("you have already started your journey, run `delete` to delete your character")
		return
	}

	// Create a new character in the Users map
	ctx.Reply("you hear the captain say 'land ho!'. You've arrived at the forgotten island of Alkos, ready to start a new journey")
	data.Users[ctx.Message.Author.ID] = &data.User{
		Level: 1,
		XP:    0,
		HP:    [2]int{20, 20},
		MP:    [2]int{20, 20},
		Gold:  0, Room: "RoomGreatMarya",
		Hat:     "HatNone",
		Inv:     []*ItemQuan{{Item: "ItemCanteen", Quan: 1}},
		Arsenal: []string{"WeaponBaseballBat"},
	}
}
