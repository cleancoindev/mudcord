package main

// User represents a character
type User struct {
	Level  int
	XP     int
	Gold   int
	HP     [2]int
	Combat bool
	Room   string
	Hat    string
	Inv    []*ItemQuan
}

// RemoveItem will either remove an item from a users inventory or decrement the quantity
func (user *User) RemoveItem(index int) {
	if user.Inv[index].Quan > 1 {
		user.Inv[index].Quan--
		return
	}

	user.Inv = append(user.Inv[:index], user.Inv[index+1:]...)
}

// AddItem will either add an item to a users inventory or increment the quantity
func (user *User) AddItem(item string, quan int) {
	for _, val := range user.Inv {
		if val.Item == item {
			val.Quan++
			return
		}
	}

	user.Inv = append(user.Inv, &ItemQuan{Item: item, Quan: quan})
}

// InvCount gets the total number of items in a users inventory
func (user *User) InvCount() int {
	var count int
	for _, val := range user.Inv {
		count += val.Quan
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
