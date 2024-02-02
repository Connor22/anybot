package modules

import (
	"anybot/conf"
	"anybot/helpers"
	"slices"

	"github.com/bwmarrin/discordgo"
)

type RoleConflictMod struct {
	flag uint32
	name string `default:"RoleConflict"`
}

func (roleconflictmod *RoleConflictMod) Init(modid int) {
	roleconflictmod.flag = (1 << modid)

	return
}

func (roleconflictmod *RoleConflictMod) Name() string {
	return roleconflictmod.name
}

func (roleconflictmod *RoleConflictMod) Flag() uint32 {
	return roleconflictmod.flag
}

func (roleconflictmod *RoleConflictMod) Intents() discordgo.Intent {
	intents := *new(discordgo.Intent)

	intents |= discordgo.IntentGuildMembers

	return intents
}

func (roleconflictmod *RoleConflictMod) Enabled(serverFlags uint32) bool {
	return roleconflictmod.flag|serverFlags != 0
}

func (roleconflictmod *RoleConflictMod) OnNewMember(guildMember *discordgo.GuildMemberAdd, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}

func (roleconflictmod *RoleConflictMod) OnGuildConnect(guildConnection *discordgo.GuildCreate, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}

func (roleconflictmod *RoleConflictMod) OnGuildConnectMember(guildMember *discordgo.Member, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	joinrole, verifyrole := serverConfig.GetJoinRole(), serverConfig.GetVerifyRole()

	if joinrole == "" {
		return
	}

	// Resolve conflicting roles
	if slices.Contains(guildMember.Roles, joinrole) && slices.Contains(guildMember.Roles, verifyrole) {
		helpers.RemoveRole(discord, guildMember.GuildID, guildMember.User.ID, joinrole)
	}
}

func (roleconflictmod *RoleConflictMod) OnMemberUpdate(updatedMember *discordgo.GuildMemberUpdate, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	if len(updatedMember.Roles) == len(updatedMember.BeforeUpdate.Roles) {
		return
	}

	joinrole, verifyrole := serverConfig.GetJoinRole(), serverConfig.GetVerifyRole()

	if slices.Contains(updatedMember.Roles, joinrole) {
		if (helpers.WasAdded(updatedMember, verifyrole)) ||
			(slices.Contains(updatedMember.Roles, verifyrole) && (helpers.WasAdded(updatedMember, joinrole))) {
			helpers.RemoveRole(discord, updatedMember.GuildID, updatedMember.User.ID, joinrole)
		}
	}
}
