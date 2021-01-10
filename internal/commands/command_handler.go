package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandHandler outlines the structure for a command handler.
type CommandHandler struct {
	prefix       string
	cmdInstances []Command
	cmdMap       map[string]Command
	middlewares  []Middleware

	OnError func(err error, ctx *Context)
}

// NewCommandHandler creates a CommandHandler.
func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix:       prefix,
		cmdInstances: make([]Command, 0),
		cmdMap:       make(map[string]Command),
		middlewares:  make([]Middleware, 0),
		OnError:      func(err error, ctx *Context) {},
	}
}

// RegisterCommand registers a new command.
func (c *CommandHandler) RegisterCommand(cmd Command) {
	c.cmdInstances = append(c.cmdInstances, cmd)
	for _, invoke := range cmd.Invokes() {
		c.cmdMap[invoke] = cmd
	}
}

// RegisterMiddleware registers a new command middleware.
func (c *CommandHandler) RegisterMiddleware(mw Middleware) {
	c.middlewares = append(c.middlewares, mw)
}

// HandleMessage handles newly any incoming message.
func (c *CommandHandler) HandleMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if !strings.HasPrefix(e.Content, c.prefix) || e.Author.ID == s.State.User.ID || e.Author.Bot {
		return
	}

	// Split the message without the prefix into the args.
	split := strings.Split(e.Content[len(c.prefix):], " ")
	// Checks if a command was passed.
	if len(split) < 1 {
		return
	}

	invoke := strings.ToLower(split[0])
	args := split[1:]

	// Checks if the command exists.
	cmd, ok := c.cmdMap[invoke]
	if !ok || cmd == nil {
		return
	}

	ctx := &Context{
		Session: s,
		Args:    args,
		Handler: c,
		Message: e.Message,
	}

	for _, mw := range c.middlewares {
		next, err := mw.Exec(ctx, cmd)
		if err != nil {
			c.OnError(err, ctx)
			return
		}
		if !next {
			return
		}
	}

	if err := cmd.Exec(ctx); err != nil {
		c.OnError(err, ctx)
	}
}
