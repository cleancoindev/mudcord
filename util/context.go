package util

import "github.com/bwmarrin/discordgo"

// Context contains data related to the current context of a command from the bots perspective
type Context struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
}

// Reply is just shorthand to easily send a message in reply
func (ctx *Context) Reply(message string) {
	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, ctx.Message.Author.Mention()+" "+message)
	if err != nil {
		Logger.Println("Failed to send message: " + err.Error())
	}
}

// SendEmbed is just shorthand to easily send a message in reply
func (ctx *Context) SendEmbed(embed discordgo.MessageEmbed) {
	_, err := ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, &embed)
	if err != nil {
		Logger.Println("Failed to send embed: " + err.Error())
	}
}
