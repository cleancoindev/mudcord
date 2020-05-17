package main

import (
	"math/rand"
	"time"
)

// NPC represents a generic NPC
type NPC struct {
	Name   string
	Speak  func(self NPC) string
	Dialog []string
}

var (
	// NPCKris is a bro
	NPCKris NPC = NPC{
		Name:   "Captain Kris",
		Dialog: []string{"Ahoy matey!", "Have a good time in Alkos!", "I've had this here boat for 16 years"},
		Speak:  SpeakDefault,
	}
)

// SpeakDefault returns a random dialog of an NPC
func SpeakDefault(self NPC) string {
	rand.Seed(time.Now().Unix())
	return self.Dialog[rand.Intn(len(self.Dialog))]
}
