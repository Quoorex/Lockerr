package lockerr

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Quoorex/Lockerr/internal/commands"
	"github.com/Quoorex/Lockerr/internal/events"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Run starts the discord bot.
func Run() {
	// Load the .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	s.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuildVoiceStates |
			discordgo.IntentsGuildMessages |
			discordgo.IntentsGuildMembers)

	registerEvents(s)
	registerCommands(s, "<lockerr ")

	// Open a websocket connection to Discord and begin listening.
	err = s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	s.Close()
}

// Registers all event handlers.
func registerEvents(s *discordgo.Session) {
	s.AddHandler(events.NewReadyHandler().Handle)
	s.AddHandler(events.NewMessageHandler().Handle)
}

func registerCommands(s *discordgo.Session, prefix string) {
	cmdHandler := commands.NewCommandHandler(prefix)
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			fmt.Sprintf("Command execution failed: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
