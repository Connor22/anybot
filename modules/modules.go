package modules

import (
	"anybot/conf"

	"github.com/bwmarrin/discordgo"
)

type Module interface {
	Init(int)
	Intents() discordgo.Intent
	Enabled(uint8) bool
	Name() string
	Flag() uint8
	OnGuildConnect(*discordgo.GuildCreate, *discordgo.Session, *conf.AnyGuild)
	OnGuildConnectMember(*discordgo.Member, *discordgo.Session, *conf.AnyGuild)
	OnMemberUpdate(*discordgo.GuildMemberUpdate, *discordgo.Session, *conf.AnyGuild)
	OnNewMember(*discordgo.GuildMemberAdd, *discordgo.Session, *conf.AnyGuild)
}

var AvailableModules = [...]Module{
	new(JoinRoleMod),
	new(RoleConflictMod),
}

func InitilializeModules() []Module {
	for modid, module := range AvailableModules {
		module.Init(modid)
	}

	return AvailableModules[:]
}

func SetIntents(discord *discordgo.Session) {
	for _, module := range AvailableModules {
		discord.Identify.Intents |= module.Intents()
	}
}
