package action

import (
	"github.com/tteeoo/mudcord/util"
)

// Action is used to define a generic action
type Action struct {
	Display string
	Fn      func(ctx *util.Context)
}
