package handlers

import (
	"anybot/helpers"
	"slices"

	"github.com/bwmarrin/discordgo"
)

func asyncCheckUser(guildMember *discordgo.Member, discord *discordgo.Session, joinrole string, verifyrole string) {
	if len(guildMember.Roles) < 1 {
		helpers.AddRole(discord, guildMember.GuildID, guildMember.User.ID, joinrole)
	} else if slices.Contains(guildMember.Roles, joinrole) {
		if slices.Contains(guildMember.Roles, verifyrole) {
			helpers.RemoveRole(discord, guildMember.GuildID, guildMember.User.ID, joinrole)
		}
	}
}

func onConnectConflictHandler(discord *discordgo.Session, newConnect *discordgo.GuildCreate) {
	joinrole, verifyrole := backend.GetJoinRole(newConnect.Guild.ID), backend.GetVerifyRole(newConnect.Guild.ID)

	if joinrole == "" {
		return
	}

	for member := range newConnect.Members {
		go asyncCheckUser(newConnect.Members[member], discord, joinrole, verifyrole)
	}
}
