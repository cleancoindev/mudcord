package util

import (
	"os"
)

// NoneDialog is generic text to print if a user deoes not have a character
const NoneDialog = "you do not have a character, run `start` to start your journey"

// NoneCombat is generic text to print if a user is trying to do something they cannot do in combat
const NoneCombat = "you cannot do that in combat"

// Colors contains color names and their hex value in decimal (for use in embeds)
var Colors = map[string]int{
	"red":   13382400,
	"blue":  26316,
	"green": 52326,
	"grey":  6710886,
	"black": 1118481,
	"brown": 9127187,
	"white": 16777214,
}

// CheckFatal checks if there is a fatal error, and exits accordingly
func CheckFatal(err error) {
	if err != nil {
		Logger.Fatalln(err)
		os.Exit(1)
	}
}

// CheckDB checks DB errors
// returns true if err != nil
func CheckDB(err error, ctx *Context) bool {
	if err != nil {
		Logger.Println("DB error: " + ctx.Message.Author.ID + ": " + err.Error())
		ctx.Reply("a database error occurred, please report this to <@258771223473815553>: ```" + err.Error() + "```")
		return true
	}

	return false
}

// InvalidChar will alert the user that their character is invalid
func InvalidChar(item string, ctx *Context) {
	Logger.Println("Invalid character: " + ctx.Message.Author.ID + ": " + item)
	ctx.Reply("your character has an invalid value, please report this to <@258771223473815553>: ```" + item + "```")
}
