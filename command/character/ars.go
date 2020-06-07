package character

import (
	"strconv"
	"strings"
	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/item/weapon"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const ArsHelp = "ars [weapon#]; Shows the weapons in your arsenal, optionally provide a weapon# to see info"

func Ars(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)
	currentRoom := room.Rooms[user.Room]

	// Send message if empty
	if len(user.Arsenal) <= 0 {
		ctx.Reply("your weapons arsenal is empty")
		return
	}

	inspect := true

	// Get item number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		inspect = false
	}
	num--

	// return if item number does not exist
	if num <= -1 || len(user.Arsenal) <= num {
		inspect = false
	}

	// Collect and send the data
	var embed discordgo.MessageEmbed
	if inspect {
		currentItem := item.Items[user.Arsenal[num]]
		embed = discordgo.MessageEmbed{
			Title:       currentItem.Display(),
			Description: currentItem.Desc(),
			Color:       util.Colors[currentRoom.Color],
			Fields:      currentItem.Inspect(),
			Author:      &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
		}
	} else {
		var items string
		for i, val := range user.Arsenal {
			weap := item.Items[val].(weapon.Weapon)
			items += "**" + strconv.Itoa(i+1) + ".** " + weap.Display() + "\n"
		}

		embed = discordgo.MessageEmbed{
			Title:  "Arsenal",
			Color:  util.Colors[currentRoom.Color],
			Fields: []*discordgo.MessageEmbedField{{Name: strconv.Itoa(len(user.Arsenal)) + "/3 weapons", Value: items, Inline: false}},
			Author: &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
		}
	}

	ctx.SendEmbed(embed)
}
