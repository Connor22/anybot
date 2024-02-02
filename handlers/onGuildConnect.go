package handlers

import (
	"anybot/conf"
	"anybot/modules"
	"anybot/storage"

	"github.com/bwmarrin/discordgo"
)

func asyncCheckUser(guildMember *discordgo.Member, discord *discordgo.Session, serverConfig *conf.AnyGuild, modules []modules.Module) {
	for modid, module := range modules {
		if serverConfig.Flags|(uint8(1)<<modid) != 0 {
			module.OnGuildConnectMember(guildMember, discord, serverConfig)
		}
	}
}

func onGuildConnectHandler(discord *discordgo.Session, newConnect *discordgo.GuildCreate) {
	// Start a goroutine for every member to perform initial/recovery checks
	// e.g. apply joinroles, handle conflicts, etc.
	cache := storage.GetCache()
	serverConfig := cache.GetGuild(discord, newConnect.Guild.ID)
	modules := cache.Modules

	for _, module := range modules {
		if module.Enabled(serverConfig.Flags) {
			module.OnGuildConnect(newConnect, discord, serverConfig)
		}
	}

	for _, member := range newConnect.Members {
		go asyncCheckUser(member, discord, serverConfig, modules)
	}
}
