package command

import (
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
	"github.com/tteeoo/mudcord/command/misc"
	"github.com/tteeoo/mudcord/command/option"
	"github.com/tteeoo/mudcord/command/character"
)

// Command represents a command
type Command struct {
	Exec                func(*util.Context)
	Help                string
	NoCombat, MustStart bool
}

// Commands contains all the possible commands
var Commands = map[string]Command{
	"ping": {
		Exec:      misc.Ping,
		Help:      misc.PingHelp,
		MustStart: false,
		NoCombat:  false,
	},
	"ops": {
		Exec:      option.Ops,
		Help:      option.OpsHelp,
		MustStart: true,
		NoCombat:  false,
	},
	"start": {
		Exec:      character.Start,
		Help:      character.StartHelp,
		NoCombat:  false,
	},
}

// Run will ensure the user has a started character if required
func (cmd *Command) Run(ctx *util.Context) {
	if cmd.Help == character.StartHelp {
		if db.CheckStarted(ctx.Message.Author.ID) {
			ctx.Reply("you have already started your journey, run `delete` to delete your character")
			return
		}
	}

	if cmd.MustStart {
		if !db.CheckStarted(ctx.Message.Author.ID) {
			ctx.Reply(util.NoneDialog)
			return
		}
	}

	if cmd.NoCombat {
		user, _ := db.GetUser(ctx.Message.Author.ID)
		if user.Combat {
			ctx.Reply(util.NoneCombat)
			return
		}
	}

	cmd.Exec(ctx)
}
