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
}

var Slime = Enemy{
	Name: "Slime",
	Desc: "A green, slippery, ball of hate",
	HP:   [2]int{10, 10},
	MP:   [2]int{0, 0},
	Attacks: []Attack{
		{
			Name:     "Body slam",
			Action:   "throws its blubbery mass of gluton in your direction",
			Damage:   5,
			Accuracy: 6,
			Crit:     5,
			Speed:    3,
			MP:       0,
		},
	},
}
