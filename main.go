package main

import (
	"anybot/handlers"
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

	// Database access
	db storage.Storage
)

func init() {
	flag.Parse()
	db = *storage.InitDB()
}

func main() {
	session := initBot()

	defer session.Close()
	defer db.Backend.Close()

	handlers.Init(session, db)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}

func initBot() *discordgo.Session {
	session, _ := discordgo.New("Bot " + *BotToken)

	// Intent Bits
	session.Identify.Intents |= discordgo.IntentAutoModerationExecution
	session.Identify.Intents |= discordgo.IntentGuildPresences
	session.Identify.Intents |= discordgo.IntentGuildMembers

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
