package db

import ()

// User represents a character
type User struct {
	ID, Room, Hat   string
	Level, Gold, XP int
	HP, MP          [2]int
	Combat          bool
	Inv             []*itemQuan
	Arsenal         []string
}
