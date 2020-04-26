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

	_, exists := Users[m.Author.ID]
	if exists {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you have already started your journey, run `.delete` to delete your character")
	} else {
		//args := strings.Split(m.Content, " ")[1:]
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" starting your journey")
		Users[m.Author.ID] = map[string]string{"level": "0", "room": "RoomSpawn"}
	}

}

// CommandDelete is used to delete a players data
func CommandDelete(s *discordgo.Session, m *discordgo.MessageCreate) {

	_, exists := Users[m.Author.ID]
	if exists {
		delete(Users, m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" successfully deleted your character, run `.start` to start a new one")
	} else {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you do not have a character to delete")
	}

}
