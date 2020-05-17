package command

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const ArsHelp = "ars <weapon#>; moves a weapon from your inventory to your weapons arsenal"

func Ars(ctx *util.Context) {

	// return and send message if character is not started
	if !data.CheckStarted(ctx.Message.Author.ID) {
		ctx.Reply(util.NoneDialog)
		return
	}

	user := data.Users[ctx.Message.Author.ID]
	room := data.Rooms[user.Room]

	// Cannot use in combat
	if user.Combat {
		ctx.Reply(util.NoneCombat)
		return
	}

	// Send message if empty
	if len(user.Arsenal) <= 0 {
		ctx.Reply("your weapons arsenal is empty")
		return

	}

	// Collect and send the data
	var items string
	for i, val := range user.Arsenal {
		items += "**" + strconv.Itoa(i+1) + ".** " + Items[val].Display + "\n"
	}

	embed := discordgo.MessageEmbed{
		Title:  "Arsenal",
		Color:  util.Colors[room.Color],
		Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: strconv.Itoa(len(user.Arsenal)) + "/3 weapons", Value: items, Inline: false}},
		Author: &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
	}

	ctx.SendEmbed(embed)
}
