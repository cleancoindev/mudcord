package command

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const ItemHelp = "item <item#>; displays more information about an item in your inventory"

func Item(ctx *data.Context) {

	// return and send message if character is not started
	if !data.CheckStarted(ctx.Message.Author.ID) {
		ctx.Reply(util.NoneDialog)
		return
	}

	user := data.Users[m.Author.ID]
	room := data.Rooms[user.Room]

	// Get item number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		ctx.Reply("that item does not exist")
		return
	}
	num--

	// return if item number does not exist
	if num <= -1 || len(user.Inv) <= num {
		ctx.Reply("that item does not exist")
		return
	}

	// Make the fields
	item := Items[user.Inv[num].Item]

	fields := []*discordgo.MessageEmbedField{
		{Name: "Type", Value: item.Type(), Inline: true},
		{Name: "Usable", Value: strconv.FormatBool(item.Usable), Inline: true},
		{Name: "Amount", Value: strconv.Itoa(user.Inv[num].Quan), Inline: true},
		{Name: "Combat usable", Value: strconv.FormatBool(item.CombatUsable), Inline: true},
	}

	// Collect and send the data
	embed := discordgo.MessageEmbed{
		Title:       item.Display,
		Description: item.Desc,
		Color:       Colors[room.Color],
		Fields:      fields,
		Author:      &discordgo.MessageEmbedAuthor{Name: m.Author.Username, IconURL: m.Author.AvatarURL("")},
	}

	ctx.SendEmbed(embed)
}
