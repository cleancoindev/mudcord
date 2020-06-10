package room

import (
	"github.com/tteeoo/mudcord/enemy"
	"github.com/tteeoo/mudcord/room/action"
	"github.com/tteeoo/mudcord/room/npc"
	"math/rand"
	"time"
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
	RoomGreatMarya.ID:  RoomGreatMarya,
	RoomWestDocks.ID:   RoomWestDocks,
	RoomTestingPath.ID: RoomTestingPath,
}

type encounterRate struct {
	percent int
	enemies []enemy.Enemy
}

// Types maps room type strings to encounters
var types = map[string]encounterRate{
	"Town": {
		percent: 0,
	},
	"Path": {
		percent: 25,
		enemies: []enemy.Enemy{
			enemy.Slime,
			enemy.Zombie,
		},
	},
}

// GetEnemy gets a random enemy (or none) based on a room's type
func (room *Room) GetEnemy() enemy.Enemy {
	enc := types[room.Type]

	rand.Seed(time.Now().Unix())
	num := rand.Intn(100)
	if num < enc.percent {
		return enc.enemies[rand.Intn(len(enc.enemies))]
	}

	return enemy.Enemy{
		Name: "None",
	}
}
