package commands

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/Quoorex/Lockerr/internal/storage"
	"github.com/zekroTJA/shireikan"
)

// Unlock is a command temporarely locking
// someone into a voice channel.
type Unlock struct {
}

// GetInvokes returns the command invokes.
func (c *Unlock) GetInvokes() []string {
	return []string{"unlock", "ul"}
}

// GetDescription returns the commands description.
func (c *Unlock) GetDescription() string {
	return "Unlocks a user no matter whether he is temporarely or permanently locked."
}

// GetHelp returns the commands help text.
func (c *Unlock) GetHelp() string {
	return "`unlock` - unlock"
}

// GetGroup returns the commands group.
func (c *Unlock) GetGroup() string {
	return shireikan.GroupGuildAdmin
}

// GetDomainName returns the commands domain name.
func (c *Unlock) GetDomainName() string {
	return "test.fun.unlock"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Unlock) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Unlock) IsExecutableInDMChannels() bool {
	return false
}

// Exec is the commands execution handler.
func (c *Unlock) Exec(ctx shireikan.Context) error {
	tempFile := storage.TempLockedUsers{}
	user := gabs.New()
	user.Array("Users")
	user.ArrayAppend(storage.LockedUser{ID: ctx.GetUser().ID, Channel: ctx.GetMessage().ChannelID}, "Users")
	_ = tempFile.Update(user)
	return nil
}
