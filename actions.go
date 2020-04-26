package main

import (
	"github.com/bwmarrin/discordgo"
)

// Action is used to define a generic action
type Action struct {
	Display  string
	Function func(s *discordgo.Session, m *discordgo.MessageCreate)
}