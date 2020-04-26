package main

// Room represents a generic Room
type Room struct {
	Display string
	Npcs    []NPC
	Rooms   []Room
	Actions []Action
}

// RoomSpawn is the first spawn room
var RoomSpawn Room = Room{Display: "Spawn room", Npcs: []NPC{}, Rooms: []Room{}, Actions: []Action{}}
