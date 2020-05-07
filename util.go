package main

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

// Server represents a server
type Server struct {
	Prefix string
}

// CheckFatal checks if there is a fatal error, and exits accordingly
func CheckFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}

// NoneDialog is generic text to print if a user deoes not have a character
func NoneDialog(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you do not have a character, run `.start` to start your journey")
}

func writeJSON() {
	b, err := json.MarshalIndent(Users, "", "  ")
	CheckFatal(err)
	logrus.Info("serializing user data sized ", len(b), " bytes")
	ioutil.WriteFile("users.json", b, 0644)

	sb, err := json.MarshalIndent(Servers, "", "  ")
	CheckFatal(err)
	logrus.Info("serializing server data sized ", len(sb), " bytes")
	ioutil.WriteFile("servers.json", sb, 0644)

	eb, err := json.MarshalIndent(Env, "", "  ")
	CheckFatal(err)
	logrus.Info("serializing environment data sized ", len(eb), " bytes")
	ioutil.WriteFile("env.json", eb, 0644)
}

// Serializer periodically serializes files
func Serializer(serQuit chan bool) {
	var rest int
	for {
		select {
		case <-serQuit:
			writeJSON()
			logrus.Info("Serializer shutting down safely")
			return
		default:
			if rest == 0 {
				writeJSON()
				rest = 10
			} else {
				rest--
				time.Sleep(1 * time.Second)
			}
		}
	}
}
