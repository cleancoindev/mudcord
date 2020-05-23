package npc

import (
	"math/rand"
	"time"
	"github.com/tteeoo/mudcord/util"
)

// NPC represents a generic NPC
type NPC struct {
	Name   string
	Speak  func(self NPC, ctx *util.Context) string
	Dialog []string
}


// SpeakDefault returns a random dialog of an NPC
func SpeakDefault(self NPC, _ *util.Context) string {
	rand.Seed(time.Now().Unix())
	return self.Dialog[rand.Intn(len(self.Dialog))]
}
