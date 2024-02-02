package main

import (
	"anybot/handlers"
	"anybot/modules"
	"anybot/storage"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

// Command line flags
var (
	BotToken = flag.String("token", "", "Bot authorization token")
)

func init() {
	flag.Parse()
}

func main() {
	storage.InitCache()
	defer storage.CloseDB()

	session := initBot()
	defer session.Close()

	handlers.Init(session)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}

func initBot() *discordgo.Session {
	session, _ := discordgo.New("Bot " + *BotToken)

	// Set Intent Bits
	modules.SetIntents(session)

	// Handlers
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})

	// Handle Errors
	if err := session.Open(); err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	return session
}
