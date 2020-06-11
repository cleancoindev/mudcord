package enemy

import (
	"time"
	"math/rand"
)

// Attack represents an enemy's attack
type Attack struct {
	Name, Action                      string
	Damage, Accuracy, Crit, Speed, MP int
}

// Enemy represents an enemy
type Enemy struct {
	Name, Desc string
	HP, MP     [2]int
	Attacks    []Attack
	History    []string
}

// GetAttack chooses a random attack from an enemy
func (enemy *Enemy) GetAttack() Attack {
	rand.Seed(time.Now().Unix())
	return enemy.Attacks[rand.Intn(len(enemy.Attacks))]
}
