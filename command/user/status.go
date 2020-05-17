package command

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const StatusHelp = "status; displays your character's status"

func Status(ctx *util.Context) {

	// Ensure command author has not started their journey
	if !data.CheckStarted(ctx.Message.Author.ID) {
		ctx.Reply(util.NoneDialog)
		return
	}

	// Get the current user and room
	user := data.Users[ctx.Message.Author.ID]
	room := data.Rooms[user.Room]

	// Make the fields
	fields := []*discordgo.MessageEmbedField{
		{Name: "Location", Value: Rooms[user.Room].Display, Inline: false},
		{Name: "Level", Value: strconv.Itoa(user.Level), Inline: true},
		{Name: "XP", Value: strconv.Itoa(user.XP), Inline: true},
		{Name: "HP", Value: strconv.Itoa(user.HP[0]) + "/" + strconv.Itoa(user.HP[1]), Inline: true},
		{Name: "Gold", Value: strconv.Itoa(user.Gold), Inline: true},
		{Name: "Items", Value: strconv.Itoa(user.InvCount()), Inline: true},
	}

	embed := discordgo.MessageEmbed{
		Title:  "Status",
		Color:  util.Colors[room.Color],
		Fields: fields,
		Author: &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
	}
	ctx.SendEmbed(embed)
}
