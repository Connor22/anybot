package modules

import (
	"anybot/conf"
	"anybot/helpers"
	"log"
	"slices"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var threshold_in_days int = 14

type RoleConflictMod struct {
	flag uint8
	name string `default:"RoleConflict"`
}

func (roleconflictmod *RoleConflictMod) Init(modid int) {
	roleconflictmod.flag = (1 << modid)

	return
}

func (roleconflictmod *RoleConflictMod) Name() string {
	return "RoleConflict"
}

func (roleconflictmod *RoleConflictMod) Start(discord *discordgo.Session, appID string) {
	return
}

func (roleconflictmod *RoleConflictMod) Stop(discord *discordgo.Session, appID string) {
	return
}

func (roleconflictmod *RoleConflictMod) Flag() uint8 {
	return roleconflictmod.flag
}

func (roleconflictmod *RoleConflictMod) Intents() discordgo.Intent {
	intents := *new(discordgo.Intent)

	intents |= discordgo.IntentGuildMembers

	return intents
}

func (roleconflictmod *RoleConflictMod) Enabled(serverFlags uint8) bool {
	return roleconflictmod.flag&serverFlags != 0
}

func (roleconflictmod *RoleConflictMod) OnGuildConnectMember(guildMember *discordgo.Member, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	newjoinrole, joinrole, verifyrole := conf.NEWATTENDEE, serverConfig.GetJoinRole(), serverConfig.GetVerifyRole()

	if joinrole == "" {
		return
	}

	// Resolve conflicting roles
	if slices.Contains(guildMember.Roles, joinrole) && slices.Contains(guildMember.Roles, verifyrole) {
		helpers.RemoveRole(discord, guildMember.GuildID, guildMember.User.ID, joinrole)
	}

	// Update role based on age
	if slices.Contains(guildMember.Roles, newjoinrole) {
		if !checkAgeUnderThreshold(threshold_in_days, guildMember.User.ID) {
			helpers.RemoveRole(discord, guildMember.GuildID, guildMember.User.ID, conf.NEWATTENDEE)
		}
	}
}

func (roleconflictmod *RoleConflictMod) OnMemberUpdate(updatedMember *discordgo.GuildMemberUpdate, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	if updatedMember.BeforeUpdate != nil && len(updatedMember.Roles) == len(updatedMember.BeforeUpdate.Roles) {
		return
	}

	newjoinrole, joinrole, verifyrole := conf.NEWATTENDEE, serverConfig.GetJoinRole(), serverConfig.GetVerifyRole()

	// Update role based on age
	if slices.Contains(updatedMember.Roles, newjoinrole) {
		if !checkAgeUnderThreshold(threshold_in_days, updatedMember.User.ID) {
			helpers.RemoveRole(discord, updatedMember.GuildID, updatedMember.User.ID, conf.NEWATTENDEE)
		}
	}

	if slices.Contains(updatedMember.Roles, joinrole) {
		if (helpers.WasAdded(updatedMember, verifyrole)) ||
			(slices.Contains(updatedMember.Roles, verifyrole) && (helpers.WasAdded(updatedMember, joinrole))) {
			helpers.RemoveRole(discord, updatedMember.GuildID, updatedMember.User.ID, joinrole)
			if checkAgeUnderThreshold(threshold_in_days, updatedMember.User.ID) {
				helpers.AddRole(discord, updatedMember.GuildID, updatedMember.User.ID, conf.NEWATTENDEE)
				//helpers.RemoveRole(discord, updatedMember.GuildID, updatedMember.User.ID, conf.ATTENDEE)
				log.Println(updatedMember.User.Username, "'s account is younger than ", threshold_in_days, " days, applying new role")
			} else {
				log.Println(updatedMember.User.Username, "'s account is older than ", threshold_in_days, " days")
			}
		}
	}
}

func checkAgeUnderThreshold(threshold int, userID string) bool {
	currentTimestamp := time.Now().UnixMilli()
	userSnowflake, err := strconv.ParseInt(userID, 10, 64)

	if err != nil {
		log.Fatalln(err)
	}
	userCreationTimestamp := helpers.ReverseSnowflake(userSnowflake)

	return ((userCreationTimestamp) > (currentTimestamp - int64(86400000*threshold)))
}
