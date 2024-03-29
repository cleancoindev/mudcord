package option

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/item/consumable"
	"github.com/tteeoo/mudcord/item/weapon"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const OpsHelp = "ops; shows currentRoom.specific options"

func Ops(ctx *util.Context) {

	// Get the players current room
	user, err := db.GetUser(ctx.Message.Author.ID)
	if util.CheckDB(err, ctx) {
		return
	}
	currentRoom := room.Rooms[user.Room]

	// Initialize fields slice
	var fields []*discordgo.MessageEmbedField
	if !user.Combat {
		roomValue := "run `go #` to travel\n"

		// Iterate all currentRoom. linked to by the current one and add them to the currentRoom.field
		for i, v := range currentRoom.Rooms {
			roomValue += "**" + strconv.Itoa(i+1) + ".**\t" + room.Rooms[v].Display + "\n"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Rooms", Value: roomValue, Inline: false})

		// If the currentRoom.has actions, iterate them and add them as a field
		if len(currentRoom.Actions) > 0 {
			actionsValue := "run `act #` to act\n"
			for i, v := range currentRoom.Actions {
				actionsValue += "**" + strconv.Itoa(i+1) + ".**\t" + v.Display + "\n"
			}
			fields = append(fields, &discordgo.MessageEmbedField{Name: "Actions", Value: actionsValue, Inline: false})
		}

		// If the currentRoom.has npcs, iterate them and add them as a field
		if len(currentRoom.NPCs) > 0 {
			npcsValue := "run `talk #` to talk\n"
			for i, v := range currentRoom.NPCs {
				npcsValue += "**" + strconv.Itoa(i+1) + ".**\t" + v.Name + "\n"
			}
			fields = append(fields, &discordgo.MessageEmbedField{Name: "NPCs", Value: npcsValue, Inline: false})
		}
	} else {

		// Ensure user contains an enemy struct
		if len(user.Facing) < 1 {
			ctx.Reply("an error occurred, please report this to <@258771223473815553>: ```" + user.ID + " has len(user.Facing) < 1```")
			return
		}

		// Display combat info
		fields = append(fields, &discordgo.MessageEmbedField{Name: "You are in combat", Value: "**" + user.Facing[0].Name + "**: " + user.Facing[0].Desc, Inline: false})

		// Display pass command
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Pass", Value: "run `pass` to skip your turn", Inline: false})

		// Display flee command if you haven't already tried to flee
		if !user.Facing[0].Fleed {
			fields = append(fields, &discordgo.MessageEmbedField{Name: "Flee", Value: "run `flee` to try and run away", Inline: false})
		}

		// Display usable items if you have > 0
		var itemCount int
		for _, v := range user.Inv {
			currentItem := item.Items[v.ID]
			switch currentItem.(type) {
			case consumable.Consumable:
				if currentItem.(consumable.Consumable).CombatUsable {
					itemCount += v.Quan
				}
			}
		}
		if itemCount > 0 {
			fields = append(fields, &discordgo.MessageEmbedField{Name: "Items", Value: "you have " + strconv.Itoa(itemCount) + " item(s) that you can use while in combat", Inline: false})
		}

		// Display weapons that are not on cooldown
		attackValue := "run `attack #` to use a weapon\n"
		for i, v := range user.Arsenal {
			cooldown := 3 - item.Items[v].(weapon.Weapon).Speed
			used := false
			for _, name := range user.History[len(user.History)-cooldown : len(user.History)] {
				if name == item.Items[v].Display() {
					used = true
				}
			}
			if !used {
				attackValue += "**" + strconv.Itoa(i+1) + ".**\t" + item.Items[v].Display() + "\n"
			}
		}
		if attackValue != "run `attack #` to use a weapon\n" {
			fields = append(fields, &discordgo.MessageEmbedField{Name: "Weapons", Value: attackValue, Inline: false})
		}
	}

	// Send an embed containing all the fields
	embed := discordgo.MessageEmbed{
		Title:       currentRoom.Display,
		Color:       util.Colors[currentRoom.Color],
		Description: currentRoom.Desc,
		Fields:      fields,
		Author:      &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
	}

	ctx.SendEmbed(embed)
}
