package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tteeoo/mudcord/util"
	"github.com/tteeoo/mudcord/data"
)

func Prefix(ctx *util.Context) {

	// Check user permission
	server, err := ctx.Session.Guild(ctx.Message.GuildID)
	if err != nil {
		ctx.Reply"you cannot change the prefix in a direct message")
		return
	}
	if ctx.Message.Author.ID != server.OwnerID {
		ctx.Reply("only the owner of the server can change the prefix")
		return
	}

	// Get the new prefix and set it
	contentSplit := strings.Split(ctx.Message.Content, " ")
	if len(contentSplit) <= 1 {
		ctx.Reply("no prefix provided")
		return
	}
	newPrefix := contentSplit[1]
	data.Servers[ctx.Message.GuildID].Prefix = newPrefix
	ctx.Reply("set the prefix to "+newPrefix)
}
