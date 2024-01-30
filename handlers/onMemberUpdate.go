package handlers

import (
	"anybot/helpers"
	"slices"
	"sync"

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

	discordMutexes[updatedMember.User.ID].Unlock()
}

func onMemberUpdateHandler(discord *discordgo.Session, updatedMember *discordgo.GuildMemberUpdate) {
	// Check if we need to further examine this event
	if !(len(updatedMember.Roles) > len(updatedMember.BeforeUpdate.Roles)) {
		return
	}

	// Make sure nothing else is happening to this user
	discordMutexes[updatedMember.GuildID].Lock()
	mutex, mutexExists := discordMutexes[updatedMember.User.ID]
	if !mutexExists {
		discordMutexes[updatedMember.User.ID] = new(sync.Mutex)
		mutex = discordMutexes[updatedMember.User.ID]
	}
	discordMutexes[updatedMember.GuildID].Unlock()

	mutex.Lock()
	go resolveVerificationConflicts(discord, updatedMember)
}
