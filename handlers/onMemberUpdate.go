package handlers

import (
	"anybot/storage"

	"github.com/bwmarrin/discordgo"
)

func onMemberUpdateHandler(discord *discordgo.Session, updatedMember *discordgo.GuildMemberUpdate) {
	cache := storage.GetCache()
	serverConfig := cache.GetGuild(discord, updatedMember.GuildID)
	modules := cache.Modules

	for _, module := range modules {
		if module.Enabled(serverConfig.Flags) {
			module.OnMemberUpdate(updatedMember, discord, serverConfig)
		}
	}
}
