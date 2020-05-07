package main

// Room represents a generic Room
type Room struct {
	Into    string
	Display string
	Desc    string
	Color   string
	NPCs    []NPC
	Rooms   []string
	Actions []Action
}

var (
	// Colors contains color names and their hex value in decimal (for use in embeds)
	Colors map[string]int = map[string]int{
		"red":   13382400,
		"blue":  26316,
		"green": 52326,
		"grey":  6710886,
		"black": 1118481,
		"brown": 9127187,
		"white": 16777214,
	}

	// Rooms contains all the rooms
	Rooms map[string]Room = map[string]Room{
		"RoomGreatMarya": RoomGreatMarya,
		"RoomWestDocks":  RoomWestDocks,
	}

	// DefaultEnv contains the default persistent data for all rooms
	DefaultEnv = map[string]int{
		"RoomSpawnCampfire": 0,
	}

	// RoomGreatMarya is the ship you arrive on
	RoomGreatMarya Room = Room{
		Into:    "the captain welcomes you aboard, the boat's not going anywhere just yet",
		Display: "The Great Marya",
		Desc:    "A medium-sized ship used to transport tourists and travelers alike to and from Alkos",
		Color:   "blue",
		NPCs:    []NPC{NPCKris},
		Rooms:   []string{"RoomWestDocks"},
		Actions: []Action{},
	}

	// RoomWestDocks is the ship you arrive on
	RoomWestDocks Room = Room{
		Into:    "you walk onto the wobbly floating docks",
		Display: "The Western Docks",
		Desc:    "A large labyrinth of platforms connecting many docked vessels",
		Color:   "brown",
		NPCs:    []NPC{},
		Rooms:   []string{"RoomGreatMarya"},
		Actions: []Action{},
	}
)
