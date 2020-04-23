package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// Token is the bot's authentication token which is obtained via environment variable
var Token string = os.Getenv("MUDCORD_TOKEN")

func main() {

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
