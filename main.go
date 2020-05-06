package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// Token is the bot's authentication token which is obtained via environment variable
var Token string = os.Getenv("MUDCORD_TOKEN")

// Users stores all the information about users
var Users map[string]*User

// Servers stores all the information about servers
var Servers map[string]*Server

// Env stores all the persistent information about rooms
var Env map[string]int

func main() {

	// Deserialize our data
	b, err := ioutil.ReadFile("users.json")
	CheckFatal(err)
	err = json.Unmarshal(b, &Users)
	CheckFatal(err)

	sb, err := ioutil.ReadFile("servers.json")
	CheckFatal(err)
	err = json.Unmarshal(sb, &Servers)
	CheckFatal(err)

	if _, err := os.Stat("env.json"); os.IsNotExist(err) {
		Env = DefaultEnv
	} else {
		eb, err := ioutil.ReadFile("env.json")
		CheckFatal(err)
		err = json.Unmarshal(eb, &Env)
		CheckFatal(err)
	}

	// Start serialization goroutine
	serQuit := make(chan bool)
	go Serializer(serQuit)

	// Make bot
	bot, err := discordgo.New("Bot " + Token)
	CheckFatal(err)

	// Add handlers
	bot.AddHandler(Ready)
	bot.AddHandler(MessageCreate)

	// Open connection
	CheckFatal(bot.Open())
	defer bot.Close()

	// Listen for ^C or other signals to stop
	logrus.Info("mudcord starting")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	serQuit <- true
	logrus.Info("Shutting down")

}
