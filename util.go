package main

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

// CheckFatal checks if there is a fatal error, and exits accordingly
func CheckFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}

// Serializer periodically serializes files
func Serializer() {
	for {
		b, err := json.MarshalIndent(Users, "", "\t")
		CheckFatal(err)
		logrus.Info("serializing data sized ", len(b), " bytes")
		ioutil.WriteFile("users.json", b, 0644)
		time.Sleep(10 * time.Second)
	}
}

// CheckStarted checks if a user has a character
func CheckStarted(id string) bool {
	_, exists := Users[id]
	if exists {
		return true
	}
	return false
}

// NoneDialog is generic text to print if a user deoes not have a character
func NoneDialog(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you do not have a character, run `.start` to start your journey")
}

// User represents a character
type User struct {
	Level int
	XP    int
	HP    int
	Room  Room
	Hat   Item
	Inv   []Item
}
