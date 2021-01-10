package commands

// Command defines the structure for each bot command
type Command interface {
	Invokes() []string
	Description() string
	AdminRequired() bool
	Exec(ctx *Context) error
}
