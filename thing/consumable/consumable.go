package consumable

import (
	"github.com/tteeoo/mudcord/util"
	"github.com/tteeoo/mudcord/thing"
)

type Consumable struct {
	Price             int
	Desc, Display, ID string
	CombatUsable      bool
	Use               func(*util.Context)
}
