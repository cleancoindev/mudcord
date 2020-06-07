package misc

import (
	"strings"

	"github.com/tteeoo/mudcord/db"
	"github.com/tteeoo/mudcord/util"
)

const PrefixHelp = "prefix <new bot prefix>; sets the bot's prefix (must be the server owner)"

func Prefix(ctx *util.Context) {

	// Check user permission
	discordServer, err := ctx.Session.Guild(ctx.Message.GuildID)
	if err != nil {
		ctx.Reply("you cannot change the prefix in a direct message")
		return
	}
	if ctx.Message.Author.ID != discordServer.OwnerID {
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

	server, err := db.GetServer(ctx.Message.GuildID)
	if util.CheckDB(err, ctx) {
		return
	}
	server.Prefix = newPrefix
	ctx.Reply("set the prefix to " + newPrefix)
	err = db.SetServer(server)
	util.CheckDB(err, ctx)
}
