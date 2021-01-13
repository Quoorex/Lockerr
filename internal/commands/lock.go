package commands

import (
	"github.com/Quoorex/Lockerr/internal/storage"
	"github.com/zekroTJA/shireikan"
)

// Lock is a command temporarely locking
// someone into a voice channel.
type Lock struct {
}

// GetInvokes returns the command invokes.
func (c *Lock) GetInvokes() []string {
	return []string{"lock", "l", "templock"}
}

// GetDescription returns the commands description.
func (c *Lock) GetDescription() string {
	return "Locks a user into a voice channel until he disconnects."
}

// GetHelp returns the commands help text.
func (c *Lock) GetHelp() string {
	return "`lock` - lock"
}

// GetGroup returns the commands group.
func (c *Lock) GetGroup() string {
	return shireikan.GroupGuildAdmin
}

// GetDomainName returns the commands domain name.
func (c *Lock) GetDomainName() string {
	return "test.fun.lock"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Lock) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Lock) IsExecutableInDMChannels() bool {
	return false
}

// Exec is the commands execution handler.
func (c *Lock) Exec(ctx shireikan.Context) error {
	user := storage.LockedUser{ID: ctx.GetUser().ID, Channel: "797732546892005386"}
	u := storage.TempLockedUsers{Users: []storage.LockedUser{user}}
	u.Write()
	return nil
}
