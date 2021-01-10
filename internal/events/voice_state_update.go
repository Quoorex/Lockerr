package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// VoiceStateUpdateHandler holds config options.
type VoiceStateUpdateHandler struct{}

// NewVoiceStateUpdateHandler returns a new VoiceStateUpdateHandler.
func NewVoiceStateUpdateHandler() *VoiceStateUpdateHandler {
	return &VoiceStateUpdateHandler{}
}

// Handle handles the discordgo.VoiceStateUpdate event
func (h *VoiceStateUpdateHandler) Handle(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
	user, err := s.User(e.UserID)
	if err != nil {
		fmt.Printf("Could not find user with the ID: %s\n", e.UserID)
	}

	if e.ChannelID != "" {
		channel, err := s.Channel(e.ChannelID)
		if err != nil {
			fmt.Printf("Could not find a voice channel with the ID: %s\n", e.ChannelID)
		}

		fmt.Printf("User %s#%s joined the voice channel %s\n", user.Username, user.Discriminator, channel.Name)
	} else {
		fmt.Printf("User %s#%s left a voice channel\n", user.Username, user.Discriminator)
	}
}
