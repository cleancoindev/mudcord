package main

import (
	"github.com/bwmarrin/discordgo"
	"math"
	"strconv"
	"strings"
	"time"
)

// Room commands

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
			actionsValue += "**" + strconv.Itoa(i+1) + ".**\t" + v.Display + "\n"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Actions", Value: actionsValue, Inline: false})
	}

	// If the room has npcs, iterate them and add them as a field
	if len(room.NPCs) > 0 {
		npcsValue := "run `talk #` to talk\n"
		for i, v := range room.NPCs {
			npcsValue += "**" + strconv.Itoa(i+1) + ".**\t" + v.Name + "\n"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "NPCs", Value: npcsValue, Inline: false})
	}

	// Send an embed containing all the fields
	embed := discordgo.MessageEmbed{
		Title:       room.Display,
		Color:       Colors[room.Color],
		Description: room.Desc,
		Fields:      fields,
		Author:      &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")},
	}

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

// CommandTalk is used to talk to an npc
func CommandTalk(s *discordgo.Session, m *discordgo.MessageCreate) {

	// return and send message if character is not started
	if !CheckStarted(m.Author.ID) {
		NoneDialog(s, m)
		return
	}

	// Get the players current room
	room := Rooms[Users[m.Author.ID].Room]

	// Get npc number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(m.Content, " ")[len(strings.Split(m.Content, " "))-1:][0])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that npc does not exist")
		return
	}
	num--

	// return if npc number does not exist
	if num <= -1 || len(room.NPCs) <= num {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that npc does not exist")
		return
	}

	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" **"+room.NPCs[num].Name+":** "+room.NPCs[num].Speak(room.NPCs[num]))
}

// Character commands

// CommandStart is used to start a player out
func CommandStart(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ensure command author has not started their journey
	if CheckStarted(m.Author.ID) {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you have already started your journey, run `delete` to delete your character")
		return
	}

	// Create a new character in the Users map
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" starting your journey")
	Users[m.Author.ID] = &User{Level: 1, XP: 0, HP: [2]int{20, 20}, Gold: 0, Room: "RoomSpawn", Hat: "HatNone", Inv: []ItemQuan{{Item: "ItemCanteen", Quan: 1}}}
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

// CommandStatus is used to see a players hp, level, gold, etc
func CommandStatus(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ensure command author has not started their journey
	if !CheckStarted(m.Author.ID) {
		NoneDialog(s, m)
		return
	}

	// Get the current user and room
	user := Users[m.Author.ID]
	room := Rooms[user.Room]

	// Make the fields
	fields := []*discordgo.MessageEmbedField{
		{Name: "Location", Value: Rooms[user.Room].Display, Inline: false},
		{Name: "Level", Value: strconv.Itoa(user.Level), Inline: true},
		{Name: "XP", Value: strconv.Itoa(user.XP), Inline: true},
		{Name: "HP", Value: strconv.Itoa(user.HP[0]) + "/" + strconv.Itoa(user.HP[1]), Inline: true},
		{Name: "Gold", Value: strconv.Itoa(user.Gold), Inline: true},
		{Name: "Items", Value: strconv.Itoa(GetInvCount(user)), Inline: true},
	}

	embed := discordgo.MessageEmbed{
		Title:  "Status",
		Color:  Colors[room.Color],
		Fields: fields,
		Author: &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

// CommandInv is used to display the contents of a users inventory
func CommandInv(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ensure command author has not started their journey
	if !CheckStarted(m.Author.ID) {
		NoneDialog(s, m)
		return
	}

	// Get the current user and room
	user := Users[m.Author.ID]
	room := Rooms[user.Room]

	if len(user.Inv) < 1 {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" you have no items in your inventory")
		return
	}

	// Get the needed amount of pages
	// Sweet, sweet, pagination /s
	var pageCount int
	if len(user.Inv)%7 != 0 {
		pageCount = int(math.Round(float64(len(user.Inv)) / 7))
		if float64(pageCount) < float64(len(user.Inv))/7.0 {
			pageCount++
		}
	} else {
		pageCount = len(user.Inv) / 7
	}

	// Make a map of every page
	var pages = make(map[int][]ItemQuan)
	for i := 1; i <= pageCount; i++ {
		upper := i + 6
		if upper > len(user.Inv) {
			upper = len(user.Inv)
		}
		pages[i] = user.Inv[i-1 : upper]
	}

	// Get page number, default 1
	num, err := strconv.Atoi(strings.Split(m.Content, " ")[len(strings.Split(m.Content, " "))-1:][0])
	if err != nil {
		num = 1
	}

	// return if page number does not exist
	if num < 1 || pageCount < num {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that page does not exist")
		return
	}

	// Get the slice of items in specific page
	var items string
	for i, val := range pages[num] {
		items += "**" + strconv.Itoa(num*7+i-6) + ".** " + Items[val.Item].Display + " (" + strconv.Itoa(val.Quan) + ")\n"
	}

	// Collect and send the data
	embed := discordgo.MessageEmbed{
		Title:  "Inventory",
		Color:  Colors[room.Color],
		Footer: &discordgo.MessageEmbedFooter{Text: strconv.Itoa(num) + "/" + strconv.Itoa(pageCount) + " pages"},
		Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: strconv.Itoa(GetInvCount(user)) + " total items", Value: items, Inline: false}},
		Author: &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

// CommandItem gives more info about an item
func CommandItem(s *discordgo.Session, m *discordgo.MessageCreate) {

	// return and send message if character is not started
	if !CheckStarted(m.Author.ID) {
		NoneDialog(s, m)
		return
	}

	user := Users[m.Author.ID]
	room := Rooms[user.Room]

	// Get item number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(m.Content, " ")[len(strings.Split(m.Content, " "))-1:][0])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that item does not exist")
		return
	}
	num--

	// return if item number does not exist
	if num <= -1 || len(user.Inv) <= num {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that item does not exist")
		return
	}

	// Make the fields
	item := Items[user.Inv[num].Item]

	fields := []*discordgo.MessageEmbedField{
		{Name: "Type", Value: item.Type, Inline: true},
		{Name: "Usable", Value: strconv.FormatBool(item.Usable), Inline: true},
		{Name: "Amount", Value: strconv.Itoa(user.Inv[num].Quan), Inline: true},
	}

	// Collect and send the data
	embed := discordgo.MessageEmbed{
		Title:       item.Display,
		Description: item.Desc,
		Color:       Colors[room.Color],
		Fields:      fields,
		Author:      &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

// CommandUse uses an item
func CommandUse(s *discordgo.Session, m *discordgo.MessageCreate) {

	// return and send message if character is not started
	if !CheckStarted(m.Author.ID) {
		NoneDialog(s, m)
		return
	}

	user := Users[m.Author.ID]

	// Get item number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(m.Content, " ")[len(strings.Split(m.Content, " "))-1:][0])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that item does not exist")
		return
	}
	num--

	// return if item number does not exist
	if num <= -1 || len(user.Inv) <= num {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" that item does not exist")
		return
	}

	Items[user.Inv[num].Item].Use(s, m)
	UserRemoveItem(user, num)
}

// Utility commands

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

// CommandPing is just a basic ping command
func CommandPing(s *discordgo.Session, m *discordgo.MessageCreate) {
	before := time.Now()
	message, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" pong!")
	if err == nil {
		ms := time.Now().Sub(before).Milliseconds()
		s.ChannelMessageEdit(message.ChannelID, message.ID, m.Author.Mention()+" pong! **"+strconv.FormatInt(ms, 10)+"ms**")
	}
}
