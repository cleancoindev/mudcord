package command

import (
	"github.com/tteeoo/mudcord/command/room"
	"github.com/tteeoo/mudcord/command/user"
	"github.com/tteeoo/mudcord/command/util"
)

// Command represents a command
type Command struct {
	Exec func(*Context)
	Help string
}

// Commands contains all the possible commands
var Commands = map[string]Command{}
