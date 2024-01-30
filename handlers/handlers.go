package handlers

import (
	"anybot/storage"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	backend   storage.Storage
	userMutex map[string]*sync.Mutex
)

func Init(discord *discordgo.Session, db storage.Storage) {
	backend = db
	userMutex = make(map[string]*sync.Mutex)

	discord.AddHandler(onMemberUpdateHandler)

	discord.AddHandler(onGuildConnectHandler)

	discord.AddHandler(onNewMemberHandler)
}
