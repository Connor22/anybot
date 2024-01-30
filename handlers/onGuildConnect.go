package handlers

import (
	"anybot/helpers"
	"slices"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func asyncCheckUser(guildMember *discordgo.Member, discord *discordgo.Session) {
	userMutex[guildMember.User.ID] = new(sync.Mutex)
	userMutex[guildMember.User.ID].Lock()
	joinrole, verifyrole := backend.GetJoinRole(guildMember.GuildID), backend.GetVerifyRole(guildMember.GuildID)

	if joinrole == "" {
		return
	}

	// Apply missing joinrole
	if len(guildMember.Roles) < 1 {
		helpers.AddRole(discord, guildMember.GuildID, guildMember.User.ID, joinrole)
	}

	// Resolve conflicting roles
	if slices.Contains(guildMember.Roles, joinrole) && slices.Contains(guildMember.Roles, verifyrole) {
		helpers.RemoveRole(discord, guildMember.GuildID, guildMember.User.ID, joinrole)
	}

	userMutex[guildMember.User.ID].Unlock()
}

func onGuildConnectHandler(discord *discordgo.Session, newConnect *discordgo.GuildCreate) {
	for member := range newConnect.Members {
		go asyncCheckUser(newConnect.Members[member], discord)
	}
}
