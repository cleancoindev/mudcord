package character

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const HatHelp = "hat; Shows info about your equipped hat"

func Hat(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)
	currentRoom := room.Rooms[user.Room]

	// Send message if empty
	if user.Hat == "None" {
		ctx.Reply("you are not wearing a hat")
		return
	}

	// Collect and send the data
	var embed discordgo.MessageEmbed
	currentItem := item.Items[user.Hat]
	embed = discordgo.MessageEmbed{
		Title:       currentItem.Display(),
		Description: currentItem.Desc(),
		Color:       util.Colors[currentRoom.Color],
		Fields:      currentItem.Inspect(),
		Author:      &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
	}

	ctx.SendEmbed(embed)
}
