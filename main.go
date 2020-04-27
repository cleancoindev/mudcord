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

func main() {

	// Deserialize our data
	b, err := ioutil.ReadFile("users.json")
	CheckFatal(err)
	err = json.Unmarshal(b, &Users)
	CheckFatal(err)

	go Serializer()

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

}
