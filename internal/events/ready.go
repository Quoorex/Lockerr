package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// ReadyHandler holds config options required for the bot startup
type ReadyHandler struct{}

// NewReadyHandler returns a new ReadyHandler.
func NewReadyHandler() *ReadyHandler {
	return &ReadyHandler{}
}

// Handler handles the discordgo.Ready event
func (h *ReadyHandler) Handler(s *discordgo.Session, e *discordgo.Ready) {
	fmt.Printf("Logged in as: %s\n", e.User)
}
