package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func Init(discord *discordgo.Session) {
	discord.AddHandler(onMemberUpdateHandler)

	discord.AddHandler(onGuildConnectHandler)

	discord.AddHandler(onNewMemberHandler)
}
