package command

import (
	"strconv"
	"strings"

	"github.com/tteeoo/mudcord/data"
	"github.com/tteeoo/mudcord/util"
)

const ArmHelp = "arm <weapon#>; moves a weapon from your weapons arsenal to your inventory"

func Arm(ctx *data.Context) {

	// return and send message if character is not started
	if !data.CheckStarted(ctx.Message.Author.ID) {
		ctx.Reply(util.NoneDialog)
		return
	}

	user := data.Users[ctx.Message.Author.ID]

	// Cannot use in combat
	if user.Combat {
		ctx.Reply(util.NoneCombat)
		return
	}

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

	user.RemoveArs(num, ctx)
}
