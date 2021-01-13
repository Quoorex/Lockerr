package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Quoorex/Lockerr/internal/commands"
	"github.com/Quoorex/Lockerr/internal/events"
	"github.com/Quoorex/Lockerr/internal/middleware"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/zekroTJA/shireikan"
)

// Starts the discord bot.
func main() {
	// Load the .env file if it exists.
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("No bot token has been configured.")
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	s.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuildVoiceStates |
			discordgo.IntentsGuildMessages |
			discordgo.IntentsGuildMembers)

	// Register all events and commands.

	registerEvents(s)

	handler := shireikan.NewHandler(&shireikan.Config{
		GeneralPrefix:         "lockerr!",
		AllowBots:             false,
		AllowDM:               true,
		ExecuteOnEdit:         true,
		InvokeToLower:         true,
		UseDefaultHelpCommand: true,
		OnError: func(ctx shireikan.Context, typ shireikan.ErrorType, err error) {
			log.Printf("[ERR] [%d] %s", typ, err.Error())
		},
	})

	handler.RegisterMiddleware(&middleware.Permissions{})

	handler.RegisterCommand(&commands.Lock{})
	handler.RegisterCommand(&commands.Unlock{})

	handler.RegisterHandlers(s)

	// Open a websocket connection to Discord and begin listening.
	err = s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Lockerr is now running.  Press CTRL-C to exit.")
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
	s.AddHandler(events.NewVoiceStateUpdateHandler().Handle)
}
