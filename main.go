package main

import (
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/tteeoo/mudcord/util"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// Token is the bot's authentication token which is obtained via environment variable
var Token string = os.Getenv("MUDCORD_TOKEN")

func main() {

	// Make bot
	bot, err := discordgo.New("Bot " + Token)
	util.CheckFatal(err)

	// Add handlers
	bot.AddHandler(Ready)
	bot.AddHandler(MessageCreate)

	// Open connection
	util.CheckFatal(bot.Open())
	defer bot.Close()

	// Listen for ^C or other signals to stop
	logrus.Info("mudcord starting")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	logrus.Info("mudcord down")
}

func ready(s *discordgo.Session, event *discordgo.Ready) {

	logrus.Info("mudcord ready")

	guilds := len(s.State.Guilds)

	s.UpdateStatus(0, "on "+strconv.Itoa(guilds)+" servers")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// return if the message is sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the server is not is Servers and add it
	_, exists := Servers[m.GuildID]
	if !exists {
		Servers[m.GuildID] = &Server{Prefix: "."}
	}

	prefix := Servers[m.GuildID].Prefix

	for name, cmd := range command.Commands {
		if prefix+name == strings.Split(m.Content, " ")[0] {
			ctx := util.Context{
				Session: s,
				Message: m,
			}

			go cmd.Exec(&ctx)
		}
	}
}
