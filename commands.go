package main

import (
	"github.com/bwmarrin/discordgo"
)

// CommandPing is just a basic ping command
func CommandPing(s *discordgo.Session, m *discordgo.MessageCreate) {

	s.ChannelMessageSend(m.ChannelID, "pong")

}
