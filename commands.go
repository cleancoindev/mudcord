package main

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
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
		Users[m.Author.ID] = User{Level: 1, XP: 0, HP: 20, Gold: 0, Room: RoomSpawn, Hat: HatNone, Inv: []Item{}}
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

		var fields []*discordgo.MessageEmbedField
		roomsValue := "run `.go #` to travel\n"
		for i, v := range room.Rooms {
			roomsValue += "**" + strconv.Itoa(i+1) + ".**\t" + Rooms[v].Display + "\n"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Rooms", Value: roomsValue, Inline: false})
		actionsValue := "run `.act #` to act\n"
		for i, v := range room.Actions {
			roomsValue += "**" + strconv.Itoa(i+1) + ".**\t" + v.Display + "\n"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Actions", Value: actionsValue, Inline: false})
		npcsValue := "run `.talk #` to talk\n"
		for i, v := range room.NPCs {
			roomsValue += "**" + strconv.Itoa(i+1) + ".**\t" + v.Name + "\n"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "NPCs", Value: npcsValue, Inline: false})

		embed := discordgo.MessageEmbed{Title: room.Display, Color: 16711680, Fields: fields, Author: &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")}}
		s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	} else {
		NoneDialog(s, m)
	}
}
