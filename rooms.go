package main

// Room represents a generic Room
type Room struct {
	Display string
	Desc    string
	NPCs    []NPC
	Rooms   []string
	Actions []Action
}

var (
	// Rooms contains all the rooms
	Rooms map[string]Room = map[string]Room{"RoomSpawn": RoomSpawn, "RoomSpawntwo": RoomSpawntwo}

	// RoomSpawn is the first spawn room
	RoomSpawn Room = Room{Display: "Spawn room", Desc: "The first room, where you spawn.", NPCs: []NPC{}, Rooms: []string{"RoomSpawntwo"}, Actions: []Action{}}

	// RoomSpawntwo is the first spawn room
	RoomSpawntwo Room = Room{Display: "Spawn room 2", Desc: "The second room, where you spawn.", NPCs: []NPC{}, Rooms: []string{"RoomSpawn"}, Actions: []Action{}}
)
