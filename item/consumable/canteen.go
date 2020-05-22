package consumable

import (
	"strconv"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
)

var Canteen = Consumable{
	Price:        10,
	Display:      "Canteen",
	Desc:         "A shiny, refillable container; heals up to two HP",
	ID:           "ConsumableCanteen",
	CombatUsable: true,
	Use:          UseCanteen,
}

func UseCanteen(ctx *util.Context) {
	user, _ := db.GetUser(ctx.Message.Author.ID)
	healed := user.Heal(2)
	if healed == 0 {
		ctx.Reply("you are already at full health")
		return
	}

	ctx.Reply("you chug down the water inside, healing " + strconv.Itoa(healed) + " health")
	user.RemoveItem(Canteen.ID)
}
