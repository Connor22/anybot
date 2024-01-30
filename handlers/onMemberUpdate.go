package handlers

import (
	"anybot/helpers"
	"slices"

	"github.com/bwmarrin/discordgo"
)

func wasAdded(member *discordgo.GuildMemberUpdate, roleid string) bool {
	return slices.Contains(member.Roles, roleid) && !slices.Contains(member.BeforeUpdate.Roles, roleid)
}

func resolveVerificationConflicts(discord *discordgo.Session, updatedMember *discordgo.GuildMemberUpdate) {
	joinrole, verifyrole := backend.GetJoinRole(updatedMember.GuildID), backend.GetVerifyRole(updatedMember.GuildID)

	if slices.Contains(updatedMember.Roles, joinrole) {
		if (wasAdded(updatedMember, verifyrole)) ||
			(slices.Contains(updatedMember.Roles, verifyrole) && (wasAdded(updatedMember, joinrole))) {
			helpers.RemoveRole(discord, updatedMember.GuildID, updatedMember.User.ID, joinrole)
		}
	}
}

func onMemberUpdateHandler(discord *discordgo.Session, updatedMember *discordgo.GuildMemberUpdate) {
	// Check if we need to further examine this event
	if !(len(updatedMember.Roles) > len(updatedMember.BeforeUpdate.Roles)) {
		return
	}

	resolveVerificationConflicts(discord, updatedMember)
}
