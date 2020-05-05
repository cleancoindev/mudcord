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
