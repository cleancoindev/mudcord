package consumable

import (
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
	"strconv"
)

var Canteen = Consumable{
	price:        10,
	display:      "Canteen",
	desc:         "A shiny, refillable container; heals up to two HP",
	id:           "ConsumableCanteen",
	combatUsable: true,
	consume:          UseCanteen,
}

func UseCanteen(ctx *util.Context) {
	user, _ := db.GetUser(ctx.Message.Author.ID)
	healed := user.Heal(2)
	if healed == 0 {
		ctx.Reply("you are already at full health")
		return
	}

	ctx.Reply("you chug down the water inside, healing " + strconv.Itoa(healed) + " health")
	user.RemoveItem("Canteen")
	db.SetUser(user)
}
