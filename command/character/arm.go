package character

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/item/weapon"
)

const ArmHelp = "arm <weapon#>; moves a weapon from your weapons arsenal to your inventory"

func Arm(ctx *util.Context) {

	user, _ := db.GetUser(ctx.Message.Author.ID)

	// Get the players arsenal
	ars := user.Arsenal

	// Get action number from message and return if it is not a number
	num, err := strconv.Atoi(strings.Split(ctx.Message.Content, " ")[len(strings.Split(ctx.Message.Content, " "))-1:][0])
	if err != nil {
		ctx.Reply("that weapon does not exist")
		return
	}
	num--

	// return if action number does not exist
	if num <= -1 || len(ars) <= num {
		ctx.Reply("that weapon does not exist")
		return
	}

	currentWeapon := user.Arsenal[num]
	user.RemoveArs(currentWeapon)
	user.AddItem(currentWeapon, 1)
	ctx.Reply("moved **" + item.Items[currentWeapon].(weapon.Weapon).Display + "** from your weapons arsenal to your inventory")
	db.SetUser(user)
}
