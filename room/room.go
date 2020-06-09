package room

import (
	"github.com/tteeoo/mudcord/enemy"
	"github.com/tteeoo/mudcord/room/action"
	"github.com/tteeoo/mudcord/room/npc"
)

// Room represents a generic Room
type Room struct {
	ID, Into, Display, Desc, Color, Type string
	NPCs                                 []npc.NPC
	Rooms                                []string
	Actions                              []action.Action
}

// Rooms contains all the rooms
var Rooms = map[string]Room{
	RoomGreatMarya.ID: RoomGreatMarya,
	RoomWestDocks.ID:  RoomWestDocks,
}

type enemyRate struct {
	percent int
	monster enemy.Enemy
}

type encounterRate struct {
	mainPercent int
	enemies     []enemyRate
}

// Types maps room type strings to encounters
var Types = map[string]encounterRate{
	"Town": {
		mainPercent: 0,
		enemies:     []enemyRate{},
	},
}
