package consumable

import (
	"github.com/tteeoo/mudcord/util"
)

type Consumable struct {
	Price             int
	Desc, Display, ID string
	CombatUsable      bool
	Use               func(*util.Context)
}
