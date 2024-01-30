package handlers

import (
	"anybot/storage"

	"github.com/bwmarrin/discordgo"
)

var (
	backend storage.Storage
)

func Init(discord *discordgo.Session, db storage.Storage) {
	backend = db

	discord.AddHandler(onMemberUpdateHandler)

	discord.AddHandler(onGuildConnectHandler)

	discord.AddHandler(onNewMemberHandler)
}
