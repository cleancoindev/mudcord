package command

import (
	"github.com/tteeoo/mudcord/command/character"
	"github.com/tteeoo/mudcord/command/misc"
	"github.com/tteeoo/mudcord/command/option"
	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/item"
	"github.com/tteeoo/mudcord/room"
	"github.com/tteeoo/mudcord/util"
)

// Command represents a command
type Command struct {
	Exec                func(*util.Context)
	Help                string
	NoCombat, MustStart bool
}

// Commands contains all the possible commands
var Commands = map[string]Command{
	// Misc. utility commands
	"ping": {
		Exec:      misc.Ping,
		Help:      misc.PingHelp,
		MustStart: false,
		NoCombat:  false,
	},
	"prefix": {
		Exec:      misc.Prefix,
		Help:      misc.PrefixHelp,
		MustStart: false,
		NoCombat:  false,
	},
	"about": {
		Exec:      misc.About,
		Help:      misc.AboutHelp,
		MustStart: false,
		NoCombat:  false,
	},

	// Room based option commands
	"ops": {
		Exec:      option.Ops,
		Help:      option.OpsHelp,
		MustStart: true,
		NoCombat:  false,
	},
	"act": {
		Exec:      option.Act,
		Help:      option.ActHelp,
		MustStart: true,
		NoCombat:  true,
	},
	"talk": {
		Exec:      option.Talk,
		Help:      option.TalkHelp,
		MustStart: true,
		NoCombat:  true,
	},
	"go": {
		Exec:      option.Go,
		Help:      option.GoHelp,
		MustStart: true,
		NoCombat:  true,
	},

	// Character based commands
	"start": {
		Exec:     character.Start,
		Help:     character.StartHelp,
		NoCombat: false,
	},
	"hat": {
		Exec:      character.Hat,
		Help:      character.HatHelp,
		MustStart: true,
		NoCombat:  false,
	},
	"hatrm": {
		Exec:      character.Hatrm,
		Help:      character.HatrmHelp,
		MustStart: true,
		NoCombat:  true,
	},
	"arm": {
		Exec:      character.Arm,
		Help:      character.ArmHelp,
		NoCombat:  true,
		MustStart: true,
	},
	"ars": {
		Exec:      character.Ars,
		Help:      character.ArsHelp,
		NoCombat:  false,
		MustStart: true,
	},
	"delete": {
		Exec:      character.Delete,
		Help:      character.DeleteHelp,
		NoCombat:  false,
		MustStart: true,
	},
	"inv": {
		Exec:      character.Inv,
		Help:      character.InvHelp,
		NoCombat:  false,
		MustStart: true,
	},
	"item": {
		Exec:      character.Item,
		Help:      character.ItemHelp,
		NoCombat:  false,
		MustStart: true,
	},
	"status": {
		Exec:      character.Status,
		Help:      character.StatusHelp,
		NoCombat:  false,
		MustStart: true,
	},
	"use": {
		Exec:      character.Use,
		Help:      character.UseHelp,
		NoCombat:  false,
		MustStart: true,
	},
	"trash": {
		Exec:      character.Trash,
		Help:      character.TrashHelp,
		NoCombat:  false,
		MustStart: true,
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

		// Validate character
		user, err := db.GetUser(ctx.Message.Author.ID)
		if util.CheckDB(err, ctx) {
			return
		}

		_, ok := item.Items[user.Hat]
		if !ok {
			util.InvalidChar(user.Hat, ctx)
			return
		}

		_, ok = room.Rooms[user.Room]
		if !ok {
			util.InvalidChar(user.Room, ctx)
			return
		}

		for _, value := range user.Inv {
			_, ok := item.Items[value.ID]
			if !ok {
				util.InvalidChar(value.ID, ctx)
				return
			}
		}

		for _, value := range user.Arsenal {
			_, ok := item.Items[value]
			if !ok {
				util.InvalidChar(value, ctx)
				return
			}
		}
	}

	if cmd.NoCombat {
		user, err := db.GetUser(ctx.Message.Author.ID)
		if util.CheckDB(err, ctx) {
			return
		}

		if user.Combat {
			ctx.Reply(util.NoneCombat)
			return
		}
	}

	cmd.Exec(ctx)
}
