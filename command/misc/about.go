package misc

import (
	"github.com/bwmarrin/discordgo"

	"github.com/tteeoo/mudcord/util"
)

const AboutHelp = "about; sends a brief message explaining more about mudcord"

func About(ctx *util.Context) {

	fields := []*discordgo.MessageEmbedField{
		{Name: "Source", Value: "mudcord is open source on [GitHub](https://github.com/tteeoo/mudcord) and licensed under the BSD license, contributions are welcome", Inline: false},
		{Name: "Credits", Value: "<@258771223473815553> (Theo Henson) - Programmer, game design\n<@546851070224105493> (Dylan Combra) - Game design", Inline: false},
	}

	embed := discordgo.MessageEmbed{
		Title:       "About",
		Description: "mudcord is a text-based mmorpg in Discord that is inspired by the multi-user dungeons of old",
		Color:       util.Colors["white"],
		Fields:      fields,
		Author:      &discordgo.MessageEmbedAuthor{Name: ctx.Session.State.Ready.User.Username, IconURL: ctx.Session.State.Ready.User.AvatarURL("")},
	}

	ctx.SendEmbed(embed)
}
