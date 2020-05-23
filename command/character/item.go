package character

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const ItemHelp = "item <item#>; displays more information about an item in your inventory"

func Item(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)
	currentRoom := room.Rooms[user.Room]

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

	// Collect and send the data
	currentItem := item.Items[user.Inv[num].ID]

	embed := discordgo.MessageEmbed{
		Title:       currentItem.Display(),
		Description: currentItem.Desc(),
		Color:       util.Colors[currentRoom.Color],
		Fields:      currentItem.Inspect(),
		Author:      &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
	}

	ctx.SendEmbed(embed)
}
