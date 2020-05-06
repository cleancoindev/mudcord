package main

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Colors contains color names and their hex value in decimal (for use in embeds)
var Colors map[string]int = map[string]int{"red": 13382400, "blue": 26316, "green": 52326, "grey": 6710886, "black": 1118481, "brown": 9127187}

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
		logrus.Info("serializing user data sized ", len(b), " bytes")
		ioutil.WriteFile("users.json", b, 0644)

		sb, err := json.MarshalIndent(Servers, "", "\t")
		CheckFatal(err)
		logrus.Info("serializing server data sized ", len(sb), " bytes")
		ioutil.WriteFile("servers.json", sb, 0644)

		eb, err := json.MarshalIndent(Env, "", "\t")
		CheckFatal(err)
		logrus.Info("serializing environment data sized ", len(eb), " bytes")
		ioutil.WriteFile("env.json", eb, 0644)

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

// DefaultSpeak returns a random dialog of an NPC
func DefaultSpeak(self NPC) string {
	rand.Seed(time.Now().Unix())
	return self.Dialog[rand.Intn(len(self.Dialog))]
}

// NoUse is for items that cnnot be used
func NoUse(item Item, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" "+item.Display+" cannot be used")
}

// User represents a character
type User struct {
	Level int
	XP    int
	Gold  int
	HP    [2]int
	Room  string
	Hat   string
	Inv   []string
}

// Server represents a server
type Server struct {
	Prefix string
}
