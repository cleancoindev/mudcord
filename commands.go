package main

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

// CommandPing is just a basic ping command
func CommandPing(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "pong")
}

// CommandStart is used to start a player out
func CommandStart(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ensure command author has not started their journey
	if CheckStarted(m.Author.ID) {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you have already started your journey, run `delete` to delete your character")
		return
	}

	// Create a new character in the Users map
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" starting your journey")
	Users[m.Author.ID] = &User{Level: 1, XP: 0, HP: [2]int{20, 20}, Gold: 0, Room: "RoomSpawn", Hat: "HatNone", Inv: []string{}}
}

// CommandDelete is used to delete a players data
func CommandDelete(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ensure command author has started their journey
	if !CheckStarted(m.Author.ID) {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you do not have a character to delete")
		return
	}

	// Delete the authors info from the Users map
	delete(Users, m.Author.ID)
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" successfully deleted your character, run `start` to start a new one")
}

// CommandOps is used display options per room
func CommandOps(s *discordgo.Session, m *discordgo.MessageCreate) {

	// return and send message if character is not started
	if !CheckStarted(m.Author.ID) {
		NoneDialog(s, m)
		return
	}

	// Get the players current room
	room := Rooms[Users[m.Author.ID].Room]

	// Initialize fields slice
	var fields []*discordgo.MessageEmbedField
	roomsValue := "run `go #` to travel\n"

	// Iterate all rooms linked to by the current one and add them to the room field
	for i, v := range room.Rooms {
		roomsValue += "**" + strconv.Itoa(i+1) + ".**\t" + Rooms[v].Display + "\n"
	}
	fields = append(fields, &discordgo.MessageEmbedField{Name: "Rooms", Value: roomsValue, Inline: false})

	// If the room has actions, iterate them and add them as a field
	if len(room.Actions) > 0 {
		actionsValue := "run `act #` to act\n"
		for i, v := range room.Actions {
			roomsValue += "**" + strconv.Itoa(i+1) + ".**\t" + v.Display + "\n"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Actions", Value: actionsValue, Inline: false})
	}

	// If the room has npcs, iterate them and add them as a field
	if len(room.NPCs) > 0 {
		npcsValue := "run `talk #` to talk\n"
		for i, v := range room.NPCs {
			roomsValue += "**" + strconv.Itoa(i+1) + ".**\t" + v.Name + "\n"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "NPCs", Value: npcsValue, Inline: false})
	}

	// Send an embed containing all the fields
	embed := discordgo.MessageEmbed{Title: room.Display, Color: Colors[room.Color], Description: room.Desc, Fields: fields, Author: &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")}}
	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

// CommandGo is used to travel to a new room
func CommandGo(s *discordgo.Session, m *discordgo.MessageCreate) {

	// return and send message if character is not started
	if !CheckStarted(m.Author.ID) {
		NoneDialog(s, m)
		return
	}

	// Get the players current room
	room := Rooms[Users[m.Author.ID].Room]

	// Get room number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(m.Content, " ")[len(strings.Split(m.Content, " "))-1:][0])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that room does not exist")
		return
	}
	num--

	// return if room number does not exist and change player room
	if num <= -1 || len(room.Rooms) <= num {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that room does not exist")
		return
	}
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" "+Rooms[room.Rooms[num]].Into)
	Users[m.Author.ID].Room = room.Rooms[num]
}

// CommandAct is used to do an action
func CommandAct(s *discordgo.Session, m *discordgo.MessageCreate) {

	// return and send message if character is not started
	if !CheckStarted(m.Author.ID) {
		NoneDialog(s, m)
		return
	}

	// Get the players current room
	room := Rooms[Users[m.Author.ID].Room]

	// Get action number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(m.Content, " ")[len(strings.Split(m.Content, " "))-1:][0])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that action does not exist")
		return
	}
	num--

	// return if action number does not exist
	if num <= -1 || len(room.Actions) <= num {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that action does not exist")
		return
	}
	room.Actions[num].Fn(s, m)
}

// CommandPrefix changes the bots prefix if you have the permission
func CommandPrefix(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Check user permission
	server, err := s.Guild(m.GuildID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you cannot change the prefix in a direct message")
		return
	}
	if m.Author.ID != server.OwnerID {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" only the owner of the server can change the prefix")
		return
	}

	// Get the new prefix and set it
	contentSplit := strings.Split(m.Content, " ")
	if len(contentSplit) <= 1 {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" no prefix provided")
		return
	}
	newPrefix := contentSplit[1]
	Servers[m.GuildID].Prefix = newPrefix
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" set the prefix to "+newPrefix)

}
