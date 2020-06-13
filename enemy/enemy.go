package enemy

import (
	"math/rand"
	"time"
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
	Fleed      bool
	Flee       int
	History    []string
}

// GetAttack chooses a random attack from an enemy, taking into
func (enemy *Enemy) GetAttack() Attack {
	var checked []Attack
	for len(checked) == len(enemy.Attacks) {
		rand.Seed(time.Now().Unix())
		attack := enemy.Attacks[rand.Intn(len(enemy.Attacks))]
		checked = append(checked, attack)
		cooldown := 3 - attack.Speed
		used := false
		for _, name := range enemy.History[len(enemy.History)-cooldown : len(enemy.History)] {
			if name == attack.Name {
				used = true
			}
		}
		if !used {
			return attack
		}
	}
	return Attack{
		Name: "None",
	}
}
