package handlers

import (
	"anybot/storage"

	"github.com/bwmarrin/discordgo"
)

func onNewMemberHandler(discord *discordgo.Session, newMember *discordgo.GuildMemberAdd) {
	cache := storage.GetCache()
	serverConfig := cache.GetGuild(discord, newMember.GuildID)
	modules := cache.Modules

	for _, module := range modules {
		if module.Enabled(serverConfig.Flags) {
			module.OnNewMember(newMember, discord, serverConfig)
		}
	}
}
