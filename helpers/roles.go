package helpers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RemoveRole(discord *discordgo.Session, gid string, uid string, roleid string) {
	err := discord.GuildMemberRoleRemove(gid, uid, roleid)
	if err != nil {
		log.Println(err)
	} else {
		role, _ := discord.State.Role(gid, roleid)
		user, _ := discord.State.Member(gid, uid)
		log.Printf("%s was removed from user %s (%s)\n", role.Name, user.User.Username, uid)
	}
}

func AddRole(discord *discordgo.Session, gid string, uid string, roleid string) {
	err := discord.GuildMemberRoleAdd(gid, uid, roleid)
	if err != nil {
		log.Println(err)
	} else {
		role, _ := discord.State.Role(gid, roleid)
		user, _ := discord.State.Member(gid, uid)
		log.Printf("%s was added to user %s (%s)\n", role.Name, user.User.Username, uid)
	}
}
