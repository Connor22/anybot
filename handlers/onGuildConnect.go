package handlers

import (
	"anybot/conf"
	mod "anybot/modules"
	"anybot/storage"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func asyncCheckUser(guildMember *discordgo.Member, discord *discordgo.Session, serverConfig *conf.AnyGuild, modules []mod.Module) {
	for _, module := range modules {
		connectMod, validModule := module.(mod.GuildConnectModule)
		if validModule && connectMod.Enabled(serverConfig.Flags) {
			connectMod.OnGuildConnectMember(guildMember, discord, serverConfig)
		}
	}
}

func onGuildConnectHandler(discord *discordgo.Session, newConnect *discordgo.GuildCreate) {
	cache := storage.GetCache()
	serverConfig := cache.GetGuild(discord, newConnect.Guild.ID)
	modules := cache.Modules

	// Run general GuildConnect functions
	for _, module := range modules {
		connectMod, validModule := module.(mod.GuildConnectModule)
		if validModule && connectMod.Enabled(serverConfig.Flags) {
			connectMod.OnGuildConnect(newConnect, discord, serverConfig)
		}
	}

	// Start a goroutine for every member to perform initial/recovery checks
	// e.g. apply joinroles, handle conflicts, etc.
	var asyncMemberThreads sync.WaitGroup
	asyncMemberThreads.Add(len(newConnect.Members))

	for _, member := range newConnect.Members {
		asyncMemberThreads.Add(1)
		go func(mem *discordgo.Member) {
			defer asyncMemberThreads.Done()
			asyncCheckUser(mem, discord, serverConfig, modules)
		}(member)
	}

	asyncMemberThreads.Wait()
}
