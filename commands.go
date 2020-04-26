package main

import (
	"github.com/bwmarrin/discordgo"
)

// CommandPing is just a basic ping command
func CommandPing(s *discordgo.Session, m *discordgo.MessageCreate) {

	s.ChannelMessageSend(m.ChannelID, "pong")

}

// CommandStart is used to start a player out
func CommandStart(s *discordgo.Session, m *discordgo.MessageCreate) {

	_, exists := Data[m.Author.ID]
	if exists {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you have already started your journey, run `.delete` to delete your character")
	} else {
		//args := strings.Split(m.Content, " ")[1:]
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" starting your journey!")
		Data[m.Author.ID] = map[string]string{"level": "0", "room": "Tutorial World"}
	}

}
