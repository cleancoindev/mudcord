package command

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const OpsHelp = "ops; shows room specific options"

func Ops(ctx *util.Context) {

	// return and send message if character is not started
	if !data.CheckStarted(ctx.Message.Author.ID) {
		ctx.Reply(util.NoneDialog)
		return
	}

	// Get the players current room
	room := Rooms[data.Users[ctx.Message.Author.ID].Room]

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
		Author:      &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
	}

	ctx.SendEmbed(embed)
}
