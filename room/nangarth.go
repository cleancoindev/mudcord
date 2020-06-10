package room

import (
	"github.com/tteeoo/mudcord/room/action"
	"github.com/tteeoo/mudcord/room/npc"
)

// RoomGreatMarya is the ship you arrive on
var RoomGreatMarya = Room{
	ID:      "RoomGreatMarya",
	Into:    "the captain welcomes you aboard, the boat's not going anywhere just yet",
	Display: "The Great Marya",
	Desc:    "A medium-sized ship used to transport tourists and travelers alike to and from Alkos",
	Color:   "blue",
	Type:    "Town",
	NPCs:    []npc.NPC{npc.NPCKris},
	Rooms:   []string{"RoomWestDocks", "RoomTestingPath"},
	Actions: []action.Action{},
}

// RoomTestingPath is a room for testing
var RoomTestingPath = Room{
	ID:      "RoomTestingPath",
	Into:    "what is this?",
	Display: "Testing Path",
	Desc:    "This room is for testing",
	Color:   "green",
	Type:    "Path",
	NPCs:    []npc.NPC{},
	Rooms:   []string{"RoomGreatMarya"},
	Actions: []action.Action{},
}

// RoomWestDocks is the ship you arrive on
var RoomWestDocks = Room{
	ID:      "RoomWestDocks",
	Into:    "you walk onto the wobbly floating docks",
	Display: "The Western Docks",
	Desc:    "A large labyrinth of platforms connecting many docked vessels",
	Color:   "brown",
	Type:    "Town",
	NPCs:    []npc.NPC{},
	Rooms:   []string{"RoomGreatMarya"},
	Actions: []action.Action{},
}
