package handlers

import (
	"anybot/helpers"
	"slices"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func asyncCheckUser(guildMember *discordgo.Member, discord *discordgo.Session) {
	discordMutexes[guildMember.User.ID] = new(sync.Mutex)
	discordMutexes[guildMember.User.ID].Lock()
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

	discordMutexes[guildMember.User.ID].Unlock()
}

func asyncHandleGuild(discord *discordgo.Session, newConnect *discordgo.GuildCreate) {
	// Set up async protection
	var checkUserWG sync.WaitGroup

	// Start a goroutine for every member to perform initial/recovery checks
	// e.g. apply joinroles, handle conflicts, etc.
	for member := range newConnect.Members {
		checkUserWG.Add(1)
		go func() {
			defer checkUserWG.Done()
			asyncCheckUser(newConnect.Members[member], discord)
		}()
	}

	// Wait for all checks on this guild to finish
	checkUserWG.Wait()
	// Release the lock so that other modifications can be made
	discordMutexes[newConnect.Guild.ID].Unlock()
}

func onGuildConnectHandler(discord *discordgo.Session, newConnect *discordgo.GuildCreate) {
	discordMutexes[newConnect.Guild.ID] = new(sync.Mutex)
	discordMutexes[newConnect.Guild.ID].Lock()
	go asyncHandleGuild(discord, newConnect)
}
