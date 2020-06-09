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
)

// Token is the bot's authentication token which is obtained via environment variable

func main() {
	defer db.Cancel()
	defer util.LogFile.Close()

	token := os.Getenv("MUDCORD_TOKEN")
	if token == "" {
		util.Logger.Fatal("MUDCORD_TOKEN not in environment")
	}

	// Make bot
	bot, err := discordgo.New("Bot " + token)
	util.CheckFatal(err)

	// Add handlers
	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)

	// Open connection
	util.CheckFatal(bot.Open())
	defer bot.Close()

	// Listen for ^C or other signals to stop
	util.Logger.Println("mudcord starting")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	util.Logger.Println("mudcord down")
}

func ready(s *discordgo.Session, event *discordgo.Ready) {

	util.Logger.Println("mudcord ready")

	guilds := len(s.State.Guilds)

	s.UpdateStatus(0, "on "+strconv.Itoa(guilds)+" servers")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer func() {
		if r := recover(); r != nil {
			util.Logger.Println("Recovered from panic in messageCreate: ", r, "; message: #"+m.ID+": "+m.Content)
		}
	}()

	// return if the message is sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Create context struct
	ctx := &util.Context{
		Session: s,
		Message: m,
	}

	// Check if the server is not is Servers and add it
	if !db.CheckServer(m.GuildID) {
		_, err := db.NewServer(m.GuildID)
		if util.CheckDB(err, ctx) {
			return
		}
	}

	server, err := db.GetServer(m.GuildID)
	if util.CheckDB(err, ctx) {
		return
	}

	prefix := server.Prefix

	// help command
	if strings.Split(m.Content, " ")[0] == prefix+"help" {
		message := "```"
		for _, cmd := range command.Commands {
			message += "" + cmd.Help + "\n"
		}
		ctx.Reply(message + "```")
		util.Logger.Println("Command: " + ctx.Message.Author.ID + ": " + prefix + "help")
	}

	// Check and run commands
	for name, cmd := range command.Commands {
		if prefix+name == strings.Split(m.Content, " ")[0] {
			cmd.Run(ctx)
			util.Logger.Println("Command: " + ctx.Message.Author.ID + ": " + ctx.Message.Content)
			return
		}
	}
}
