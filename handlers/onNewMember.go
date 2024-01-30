package handlers

import (
	"anybot/helpers"

	"github.com/bwmarrin/discordgo"
)

func onNewMemberHandler(discord *discordgo.Session, newMember *discordgo.GuildMemberAdd) {
	joinrole := backend.GetJoinRole(newMember.GuildID)

	if joinrole != "" {
		helpers.AddRole(discord, newMember.GuildID, newMember.User.ID, joinrole)
	}
}
