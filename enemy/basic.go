package enemy

var Slime = Enemy{
	Name: "Slime",
	Desc: "A green, slippery, ball of hate",
	Flee: 10,
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
		},
	},
}

var Zombie = Enemy{
	Name: "Zombie",
	Desc: "A decrepit, infected ghoul",
	Flee: 20,
	HP:   [2]int{15, 15},
	MP:   [2]int{0, 0},
	Attacks: []Attack{
		{
			Name:     "Scratch",
			Action:   "wails at you with its unruly nails",
			Damage:   4,
			Accuracy: 8,
			Crit:     2,
			Speed:    3,
		},
		{
			Name:     "Bite",
			Action:   "reaches for you with its mouth ajar",
			Damage:   7,
			Accuracy: 6,
			Crit:     10,
			Speed:    2,
		},
	},
}
