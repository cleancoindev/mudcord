package npc

import (
	"math/rand"
	"time"
)

// NPC represents a generic NPC
type NPC struct {
	Name   string
	Speak  func(NPC) string
	Dialog []string
}

// SpeakDefault returns a random dialog of an NPC
func SpeakDefault(self NPC) string {
	rand.Seed(time.Now().Unix())
	return self.Dialog[rand.Intn(len(self.Dialog))]
}
