package character

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const StatusHelp = "status; displays your character's status"

func Status(ctx *util.Context) {

	// Get the current user and room
	user, err := db.GetUser(ctx.Message.Author.ID)
	if util.CheckDB(err, ctx) {
		return
	}
	currentRoom := room.Rooms[user.Room]

	// Make the fields
	fields := []*discordgo.MessageEmbedField{
		{Name: "Location", Value: room.Rooms[user.Room].Display, Inline: false},
		{Name: "Level", Value: strconv.Itoa(user.Level), Inline: true},
		{Name: "XP", Value: strconv.Itoa(user.XP), Inline: true},
		{Name: "HP", Value: strconv.Itoa(user.HP[0]) + "/" + strconv.Itoa(user.HP[1]), Inline: true},
		{Name: "MP", Value: strconv.Itoa(user.MP[0]) + "/" + strconv.Itoa(user.MP[1]), Inline: true},
		{Name: "Gold", Value: strconv.Itoa(user.Gold), Inline: true},
		{Name: "Items", Value: strconv.Itoa(user.InvCount()), Inline: true},
	}

	embed := discordgo.MessageEmbed{
		Title:  "Status",
		Color:  util.Colors[currentRoom.Color],
		Fields: fields,
		Author: &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
	}
	ctx.SendEmbed(embed)
}
