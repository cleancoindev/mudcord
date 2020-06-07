package db

// User represents a character
type User struct {
	ID, Room, Hat   string
	Level, Gold, XP int
	HP, MP          [2]int
	Combat          bool
	Inv             []*ItemQuan
	Arsenal         []string
}

// RemoveItem will either remove an item from a users inventory or decrement the quantity
// returns false if the user does not have the item
func (user *User) RemoveItem(id string) bool {
	for index, value := range user.Inv {
		if value.ID == id {
			if user.Inv[index].Quan > 1 {
				user.Inv[index].Quan--
				return true
			}

			user.Inv = append(user.Inv[:index], user.Inv[index+1:]...)
			return true
		}
	}

	return false
}

// AddItem will either add an item to a users inventory or increment the quantity
func (user *User) AddItem(id string, quan int) {
	for _, value := range user.Inv {
		if value.ID == id {
			value.Quan++
			return
		}
	}

	user.Inv = append(user.Inv, &ItemQuan{ID: id, Quan: quan})
}

// Wear will equip a hat onto a user
// returns the id of the previous hat
func (user *User) Wear(id string) string {
	tempHat := user.Hat
	user.Hat = id
	return tempHat
}

// AddArs will attempt to add an item to a users arsenal
// returns false if arsenal is full
func (user *User) AddArs(id string) bool {
	if len(user.Arsenal) >= 3 {
		return false
	}

	user.Arsenal = append(user.Arsenal, id)
	return true
}

// RemoveArs will attempt to remove a weapon from a users arsenal
// returns false if the user does not have the weapon
func (user *User) RemoveArs(id string) bool {
	for index, value := range user.Arsenal {
		if value == id {
			user.Arsenal = append(user.Arsenal[:index], user.Arsenal[index+1:]...)
			return true
		}
	}

	return false
}

// InvCount gets the total number of items in a users inventory
func (user *User) InvCount() int {
	var count int
	for _, value := range user.Inv {
		count += value.Quan
	}

	return count
}

// Heal attempts to heal the user, taking into account max hp and returns the amount healed
func (user *User) Heal(amount int) int {
	if user.HP[0] >= user.HP[1] {
		return 0
	}

	if user.HP[0]+amount > user.HP[1] {
		user.HP[0] = user.HP[1]
		return user.HP[1] - user.HP[0]
	}

	user.HP[0] += amount
	return amount
}
