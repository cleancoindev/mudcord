package main

// NPC represents a generic NPC
type NPC struct {
	Name   string
	Speak  func(self NPC) string
	Dialog []string
}

var (
	// NPCKris is a bro
	NPCKris NPC = NPC{Name: "Kris", Dialog: []string{"BRUH", "HOWSTSAGRB:KHN"}, Speak: DefaultSpeak}
)
