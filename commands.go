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
	if CheckStarted(m.Author.ID) {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you have already started your journey, run `.delete` to delete your character")
	} else {
		//args := strings.Split(m.Content, " ")[1:]
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" starting your journey")
		Users[m.Author.ID] = User{Level: 1, XP: 0, HP: 20, Room: RoomSpawn, Hat: HatNone, Inv: []Item{}}
	}
}

// CommandDelete is used to delete a players data
func CommandDelete(s *discordgo.Session, m *discordgo.MessageCreate) {
	if CheckStarted(m.Author.ID) {
		delete(Users, m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" successfully deleted your character, run `.start` to start a new one")
	} else {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you do not have a character to delete")
	}
}

// CommandOps is used display options per room
func CommandOps(s *discordgo.Session, m *discordgo.MessageCreate) {
	if CheckStarted(m.Author.ID) {
		room := Users[m.Author.ID].Room
		embed := discordgo.MessageEmbed{Title: room.Display, Author: &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")}}
		s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	} else {
		NoneDialog(s, m)
	}
}
