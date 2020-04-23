package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"strconv"
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

	switch m.Content {
	case ".ping":
		CommandPing(s, m)
	}

}
