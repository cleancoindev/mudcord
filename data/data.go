package data

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/tteeoo/mudcord/util"
	"github.com/sirupsen/logrus"
)

// Users stores all the information about users
var Users map[string]*User

// Servers stores all the information about servers
var Servers map[string]*Server

// CheckStarted checks if a user id has a character
func CheckStarted(id string) bool {
	_, exists := Users[id]
	if exists {
		return true
	}
	return false
}

func writeJSON() {

	// Serializes the data
	b, err := json.MarshalIndent(Users, "", "  ")
	util.CheckFatal(err)
	logrus.Info("serializing user data sized ", len(b), " bytes")
	ioutil.WriteFile("users.json", b, 0644)

	sb, err := json.MarshalIndent(Servers, "", "  ")
	util.CheckFatal(err)
	logrus.Info("serializing server data sized ", len(sb), " bytes")
	ioutil.WriteFile("servers.json", sb, 0644)
}

// Serializer is to be ran as a goroutine to periodically serializes files
func Serializer(serQuit chan bool) {

	// Deserialize our data
	b, err := ioutil.ReadFile("users.json")
	util.CheckFatal(err)
	err = json.Unmarshal(b, &Users)
	util.CheckFatal(err)

	sb, err := ioutil.ReadFile("servers.json")
	util.CheckFatal(err)
	err = json.Unmarshal(sb, &Servers)
	util.CheckFatal(err)
	rest := 512

	for {
		select {
		case <-serQuit:
			writeJSON()
			logrus.Info("serializer shutting down safely")
			return
		default:
			if rest == 0 {
				writeJSON()
				rest = 2048
			} else {
				rest--
				time.Sleep(1 * time.Second)
			}
		}
	}
}
