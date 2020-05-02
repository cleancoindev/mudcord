package main

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// Ready is ran when the bot is ready
func Ready(s *discordgo.Session, event *discordgo.Ready) {

	logrus.Info("mudcord ready")

	guilds := len(s.State.Guilds)

	s.UpdateStatus(0, "in the dungeon on "+strconv.Itoa(guilds)+" servers")

}

// MessageCreate is ran when a message is created (how surprising)
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

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

	switch strings.Split(m.Content, " ")[0] {
	case prefix + "ping":
		CommandPing(s, m)
	case prefix + "start":
		CommandStart(s, m)
	case prefix + "delete":
		CommandDelete(s, m)
	case prefix + "ops":
		CommandOps(s, m)
	case prefix + "go":
		CommandGo(s, m)
	case prefix + "act":
		CommandAct(s, m)
	case prefix + "prefix":
		CommandPrefix(s, m)
	}

}
