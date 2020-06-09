package enemy

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
