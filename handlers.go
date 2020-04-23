package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	//"strconv"
)

// Ready is ran when the bot is ready
func Ready(s *discordgo.Session, event *discordgo.Ready) {

	logrus.Info("mudcord ready")

	// guilds, err := s.UserGuilds(100, "", "")
	// if err != nil {
	// 	guilds = []*discordgo.UserGuild{}
	// 	logrus.Warn(err)
	// }

	s.UpdateStatus(0, "in the dungeons")

}

// MessageCreate is ran when a message is created (how surprising)
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// return if the message is sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case ".ping":
		CommandPing(s, m)
	}

}
