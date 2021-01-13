package middleware

import (
	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/shireikan"
)

// Permissions check whether the user has the necessary
// permissions to execute the given command.
type Permissions struct {
}

// Handle is the Middlewares handler.
func (m *Permissions) Handle(cmd shireikan.Command, ctx shireikan.Context, layer shireikan.MiddlewareLayer) (bool, error) {
	if cmd.GetGroup() == shireikan.GroupGuildAdmin {
		guild, err := ctx.GetSession().Guild(ctx.GetMessage().GuildID)
		if err != nil {
			return false, err
		}
		if guild.OwnerID == ctx.GetUser().ID {
			return true, nil
		}

		roleMap := make(map[string]*discordgo.Role)
		for _, role := range guild.Roles {
			roleMap[role.ID] = role
		}

		for _, rID := range ctx.GetMember().Roles {
			if role, ok := roleMap[rID]; ok && role.Permissions&discordgo.PermissionAdministrator > 0 {
				return true, nil
			}
		}

		ctx.GetSession().ChannelMessageSend(ctx.GetChannel().ID,
			"You don't have the permission to execute this command!")
		return false, nil
	}

	// no special permissions needed
	return true, nil
}

// GetLayer returns the execution layer.
func (m *Permissions) GetLayer() shireikan.MiddlewareLayer {
	return shireikan.LayerBeforeCommand
}
