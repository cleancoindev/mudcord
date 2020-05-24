package character

import (
	"math"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/item/consumable"
	"github.com/tteeoo/mudcord/item/weapon"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

const InvHelp = "inv [page#]; displays a page of your inventory"

func Inv(ctx *util.Context) {

	// Get the current user and room
	user, _ := db.GetUser(ctx.Message.Author.ID)
	currentRoom := room.Rooms[user.Room]

	if len(user.Inv) < 1 {
		ctx.Reply("you have no items in your inventory")
		return
	}

	// Get the needed amount of pages
	// Sweet, sweet, pagination /s
	var pageCount int
	if len(user.Inv)%7 != 0 {
		pageCount = int(math.Round(float64(len(user.Inv)) / 7))
		if float64(pageCount) < float64(len(user.Inv))/7.0 {
			pageCount++
		}
	} else {
		pageCount = len(user.Inv) / 7
	}

	// Make a map of every page
	var pages = make(map[int][]*db.ItemQuan)
	for i := 1; i <= pageCount; i++ {
		upper := i + 6
		if upper > len(user.Inv) {
			upper = len(user.Inv)
		}
		pages[i] = user.Inv[i-1 : upper]
	}

	// Get page number, default 1
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		num = 1
	}

	// return if page number does not exist
	if num < 1 || pageCount < num {
		ctx.Reply("that page does not exist")
		return
	}

	// Get the slice of items in specific page
	var items string
	for i, val := range pages[num] {
		currentItem := item.Items[val.ID]
		switch currentItem.(type) {
		case consumable.Consumable:
			items += "**" + strconv.Itoa(num*7+i-6) + ".** " + currentItem.(consumable.Consumable).Display + " (" + strconv.Itoa(val.Quan) + ")\n"
		case weapon.Weapon:
			items += "**" + strconv.Itoa(num*7+i-6) + ".** " + currentItem.(weapon.Weapon).Display + " (" + strconv.Itoa(val.Quan) + ")\n"
		}
	}

	// Collect and send the data
	embed := discordgo.MessageEmbed{
		Title:  "Inventory",
		Color:  util.Colors[currentRoom.Color],
		Footer: &discordgo.MessageEmbedFooter{Text: strconv.Itoa(num) + "/" + strconv.Itoa(pageCount) + " pages"},
		Fields: []*discordgo.MessageEmbedField{{Name: strconv.Itoa(user.InvCount()) + " total items", Value: items, Inline: false}},
		Author: &discordgo.MessageEmbedAuthor{Name: ctx.Message.Author.Username, IconURL: ctx.Message.Author.AvatarURL("")},
	}

	ctx.SendEmbed(embed)
}
