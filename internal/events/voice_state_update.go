package events

import (
	"fmt"

	"github.com/Quoorex/Lockerr/internal/storage"
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
		currentChannel, err := s.Channel(e.ChannelID)
		if err != nil {
			fmt.Printf("Could not find a voice channel with the ID: %s\n", e.ChannelID)
		}

		fmt.Printf("User %s#%s joined the voice channel %s\n", user.Username, user.Discriminator, currentChannel.Name)

		// Read the list of locked users.
		fileTemp := storage.TempLockedUsers{}
		lockedUsers, _ := fileTemp.Read()
		// Check if user should be moved.
		userIsLocked, lockedUser := storage.UserIsLocked(user.ID, lockedUsers)
		if userIsLocked {
			targetChannel, err := s.Channel(lockedUser.Channel)
			if err != nil {
				println(err)
			}
			// Checks that the users isn't already in the correct channel.
			if lockedUser.Channel != currentChannel.ID {
				s.GuildMemberMove(e.GuildID, user.ID, &targetChannel.ID)
			}
		}
	} else {
		fmt.Printf("User %s#%s left a voice channel\n", user.Username, user.Discriminator)
	}
}
