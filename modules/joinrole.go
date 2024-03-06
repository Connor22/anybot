package modules

import (
	"anybot/conf"
	"anybot/helpers"

	"github.com/bwmarrin/discordgo"
)

type JoinRoleMod struct {
	flag uint8
	name string `default:"RoleConflict"`
}

func (joinmod *JoinRoleMod) Init(modid int) {
	joinmod.flag = (1 << modid)

	return
}

func (joinmod *JoinRoleMod) Name() string {
	return joinmod.name
}

func (joinmod *JoinRoleMod) Flag() uint8 {
	return joinmod.flag
}

func (joinmod *JoinRoleMod) Intents() discordgo.Intent {
	intents := *new(discordgo.Intent)

	intents |= discordgo.IntentGuilds

	return intents
}

func (joinmod *JoinRoleMod) Enabled(serverFlags uint8) bool {
	return joinmod.flag&serverFlags != 0
}

func (joinmod *JoinRoleMod) OnNewMember(guildMember *discordgo.GuildMemberAdd, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	joinrole := serverConfig.GetJoinRole()

	if joinrole != "" {
		helpers.AddRole(discord, guildMember.GuildID, guildMember.User.ID, joinrole)
	}
}

func (joinmod *JoinRoleMod) OnGuildConnectMember(guildMember *discordgo.Member, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	if len(guildMember.Roles) < 1 {
		helpers.AddRole(discord, guildMember.GuildID, guildMember.User.ID, conf.ATTENDEE)
	}
}

func (joinmod *JoinRoleMod) OnMemberUpdate(guildMember *discordgo.GuildMemberUpdate, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	joinrole := serverConfig.GetJoinRole()

	if joinrole != "" {
		helpers.AddRole(discord, guildMember.GuildID, guildMember.User.ID, joinrole)
	}
}
