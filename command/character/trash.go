package character

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/item/consumable"
	"github.com/tteeoo/mudcord/item/weapon"
	"github.com/tteeoo/mudcord/util"
)

const TrashHelp = "trash <item#> [amount]; removes an item or items from your inventory"

func Trash(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)

	// Get item number from message and return if it is not a number
	if len(strings.Split(ctx.Message.Content, " ")) <= 1 {
		ctx.Reply("provide an item to throw out")
		return
	}
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[1])
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

	// Get amount
	amount := 1
	if len(strings.Split(ctx.Message.Content, " ")) >= 3 {
		amount, err = strconv.Atoi(strings.Split(ctx.Message.Content, " ")[2])
		if err != nil || amount <= 0 {
			amount = 1
		}
	}
	if amount > user.Inv[num].Quan {
		amount = user.Inv[num].Quan
	}

	// Collect and send the data
	currentItem := item.Items[user.Inv[num].ID]

	var display string
	switch currentItem.(type) {
	case consumable.Consumable:
		display = currentItem.(consumable.Consumable).Display
	case weapon.Weapon:
		display = currentItem.(weapon.Weapon).Display
	}

	for i := 0; i < amount; i++ {
		user.RemoveItem(user.Inv[num].ID)
	}

	db.SetUser(user)

	ctx.Reply("removed " + strconv.Itoa(amount) + " **" + display + "** from your inventory")
}
