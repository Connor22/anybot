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
	Start()
}

type GuildConnectModule interface {
	Module
	OnGuildConnect(*discordgo.GuildCreate, *discordgo.Session, *conf.AnyGuild)
}

type GuildConnectMemberModule interface {
	Module
	OnGuildConnectMember(*discordgo.Member, *discordgo.Session, *conf.AnyGuild)
}

type MemberUpdateModule interface {
	Module
	OnMemberUpdate(*discordgo.GuildMemberUpdate, *discordgo.Session, *conf.AnyGuild)
}

type MemberAddModule interface {
	Module
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

func StartModules() {
	for _, module := range AvailableModules {
		module.Start()
	}
}

func SetIntents(discord *discordgo.Session) {
	discord.Identify.Intents |= discordgo.IntentsGuilds
	for _, module := range AvailableModules {
		discord.Identify.Intents |= module.Intents()
	}
}
