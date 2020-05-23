package misc

import (
	"strconv"
	"time"

	"github.com/tteeoo/mudcord/util"
)

const PingHelp = "ping; test the response time of mudcord"

func Ping(ctx *util.Context) {
	before := time.Now()
	message, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, ctx.Message.Author.Mention()+" pong!")
	if err == nil {
		ms := time.Now().Sub(before).Milliseconds()
		ctx.Session.ChannelMessageEdit(message.ChannelID, message.ID, ctx.Message.Author.Mention()+" pong! **"+strconv.FormatInt(ms, 10)+"ms**")
	}
}
