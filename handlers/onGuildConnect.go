package handlers

import (
	"anybot/helpers"
	"slices"

	"github.com/bwmarrin/discordgo"
)

func asyncCheckUser(guildMember *discordgo.Member, discord *discordgo.Session) {
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
}

func onGuildConnectHandler(discord *discordgo.Session, newConnect *discordgo.GuildCreate) {
	// Start a goroutine for every member to perform initial/recovery checks
	// e.g. apply joinroles, handle conflicts, etc.
	for member := range newConnect.Members {
		go asyncCheckUser(newConnect.Members[member], discord)
	}
}
