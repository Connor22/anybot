package handlers

import (
	"anybot/storage"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	backend        storage.Storage
	discordMutexes map[string]*sync.Mutex
)

func Init(discord *discordgo.Session, db storage.Storage) {
	backend = db
	discordMutexes = make(map[string]*sync.Mutex)

	discord.AddHandler(onMemberUpdateHandler)

	discord.AddHandler(onGuildConnectHandler)

	discord.AddHandler(onNewMemberHandler)
}
