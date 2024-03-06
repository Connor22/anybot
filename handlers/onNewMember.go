package handlers

import (
	mod "anybot/modules"
	"anybot/storage"

	"github.com/bwmarrin/discordgo"
)

func onNewMemberHandler(discord *discordgo.Session, newMember *discordgo.GuildMemberAdd) {
	cache := storage.GetCache()
	serverConfig := cache.GetGuild(discord, newMember.GuildID)
	modules := cache.Modules

	for _, module := range modules {
		newMemberMod, validModule := module.(mod.MemberAddModule)
		if validModule && newMemberMod.Enabled(serverConfig.Flags) {
			newMemberMod.OnNewMember(newMember, discord, serverConfig)
		}
	}
}
