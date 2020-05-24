package main

import (
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/tteeoo/mudcord/command"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// Token is the bot's authentication token which is obtained via environment variable
var Token string = os.Getenv("MUDCORD_TOKEN")

func main() {
	defer db.Cancel()

	// Make bot
	bot, err := discordgo.New("Bot " + Token)
	util.CheckFatal(err)

	// Add handlers
	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)

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
	defer func() {
		if r := recover(); r != nil {
			logrus.Warn("recovered from panic in messageCreate: ", r, "; message: #"+m.ID+": "+m.Content)
		}
	}()

	// return if the message is sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the server is not is Servers and add it
	if !db.CheckServer(m.GuildID) {
		db.NewServer(m.GuildID)
	}

	server, _ := db.GetServer(m.GuildID)

	prefix := server.Prefix

	// help command
	if strings.Split(m.Content, " ")[0] == prefix+"help" {
		message := "```"
		for _, cmd := range command.Commands {
			message += "" + cmd.Help + "\n"
		}
		ctx := util.Context{
			Session: s,
			Message: m,
		}
		ctx.Reply(message + "```")
	}

	for name, cmd := range command.Commands {
		if prefix+name == strings.Split(m.Content, " ")[0] {
			ctx := util.Context{
				Session: s,
				Message: m,
			}

			go cmd.Run(&ctx)
			return
		}
	}
}
