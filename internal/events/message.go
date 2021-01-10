package events

import (
	"github.com/bwmarrin/discordgo"
)

// MessageHandler holds config options
type MessageHandler struct{}

// NewMessageHandler returns anew MessageHandler.
func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

// Handle handles newly creates messages.
func (h *MessageHandler) Handle(s *discordgo.Session, e *discordgo.MessageCreate) {

}
