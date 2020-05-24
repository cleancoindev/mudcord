package character

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/item/weapon"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const ArsHelp = "ars <weapon#>; moves a weapon from your inventory to your weapons arsenal"

func Ars(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)
	currentRoom := room.Rooms[user.Room]

	// Send message if empty
	if len(user.Arsenal) <= 0 {
		ctx.Reply("your weapons arsenal is empty")
		return

	}

	// Collect and send the data
	var items string
	for i, val := range user.Arsenal {
		weap := item.Items[val].(weapon.Weapon)
		items += "**" + strconv.Itoa(i+1) + ".** " + weap.Display + "\n"
	}

	embed := discordgo.MessageEmbed{
		Title:  "Arsenal",
		Color:  util.Colors[currentRoom.Color],
		Fields: []*discordgo.MessageEmbedField{{Name: strconv.Itoa(len(user.Arsenal)) + "/3 weapons", Value: items, Inline: false}},
		Author: &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
	}

	ctx.SendEmbed(embed)
}
