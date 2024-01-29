package handlers

import (
	"anybot/helpers"
	"log"
	"slices"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	activeChecks map[string]*sync.Mutex
)

func wasAdded(member *discordgo.GuildMemberUpdate, roleid string) bool {
	return slices.Contains(member.Roles, roleid) && !slices.Contains(member.BeforeUpdate.Roles, roleid)
}

func checkConflict(discord *discordgo.Session, updatedMember *discordgo.GuildMemberUpdate) {
	activeChecks[updatedMember.User.ID].Lock()

	joinrole, verifyrole := backend.GetJoinRole(updatedMember.GuildID), backend.GetVerifyRole(updatedMember.GuildID)

	if slices.Contains(updatedMember.Roles, joinrole) {
		if (wasAdded(updatedMember, verifyrole)) ||
			(slices.Contains(updatedMember.Roles, verifyrole) && (wasAdded(updatedMember, joinrole))) {
			helpers.RemoveRole(discord, updatedMember.GuildID, updatedMember.User.ID, joinrole)
		}
	}

	activeChecks[updatedMember.User.ID].Unlock()
}

func onRoleConflictHandler(discord *discordgo.Session, updatedMember *discordgo.GuildMemberUpdate) {
	if !(len(updatedMember.Roles) > len(updatedMember.BeforeUpdate.Roles)) {
		return
	} else {
		log.Printf("User went from %d to %d roles", len(updatedMember.Roles), len(updatedMember.BeforeUpdate.Roles))
	}

	_, mutexExists := activeChecks[updatedMember.User.ID]
	if !mutexExists {
		activeChecks[updatedMember.User.ID] = new(sync.Mutex)
	}

	go checkConflict(discord, updatedMember)
}
