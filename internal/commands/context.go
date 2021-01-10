package commands

import "github.com/bwmarrin/discordgo"

// Context is the context of an invoked bot command.
type Context struct {
	Session *discordgo.Session
	Message *discordgo.Message
	Args    []string
	Handler *CommandHandler
}
