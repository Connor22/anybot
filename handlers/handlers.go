package handlers

import (
	"anybot/storage"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	backend storage.Storage
)

func Init(discord *discordgo.Session, db storage.Storage) {
	backend = db
	activeChecks = make(map[string]*sync.Mutex)

	discord.AddHandler(onRoleConflictHandler)

	discord.AddHandler(onConnectConflictHandler)

	discord.AddHandler(onNewMemberHandler)
}
