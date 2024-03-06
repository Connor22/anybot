package handlers

import (
	mod "anybot/modules"
	"anybot/storage"

	"github.com/bwmarrin/discordgo"
)

func onMemberUpdateHandler(discord *discordgo.Session, updatedMember *discordgo.GuildMemberUpdate) {
	cache := storage.GetCache()
	serverConfig := cache.GetGuild(discord, updatedMember.GuildID)
	modules := cache.Modules

	for _, module := range modules {
		memUpdateMod, validModule := module.(mod.MemberUpdateModule)
		if validModule && memUpdateMod.Enabled(serverConfig.Flags) {
			memUpdateMod.OnMemberUpdate(updatedMember, discord, serverConfig)
		}
	}
}
