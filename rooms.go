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
	Colors map[string]int = map[string]int{"red": 13382400, "blue": 26316, "green": 52326, "grey": 6710886, "black": 1118481, "brown": 9127187}

	// Rooms contains all the rooms
	Rooms map[string]Room = map[string]Room{
		"RoomSpawn":    RoomSpawn,
		"RoomSpawntwo": RoomSpawntwo,
	}

	// DefaultEnv contains the default persistent data for all rooms
	DefaultEnv = map[string]int{
		"RoomSpawnCampfire": 0,
	}

	// RoomSpawn is the first spawn room
	RoomSpawn Room = Room{Into: "You enter a familiar room (spawn)", Display: "Spawn room", Desc: "The first room, where you spawn.", Color: "brown", NPCs: []NPC{NPCKris}, Rooms: []string{"RoomSpawntwo"}, Actions: []Action{}}

	// RoomSpawntwo is the first spawn room
	RoomSpawntwo Room = Room{Into: "goint to two", Display: "Spawn room 2", Desc: "The second room, where you spawn.", Color: "black", NPCs: []NPC{}, Rooms: []string{"RoomSpawn", "RoomSpawntwo"}, Actions: []Action{}}
)
