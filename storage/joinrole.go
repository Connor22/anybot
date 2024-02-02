package storage

import "log"

func (db Cache) __todo__SetJoinRole(GuildID string, joinrole string) {
	_, err := db.connection.Exec("UPDATE guilds SET joinrole = ? WHERE guildid = ?", joinrole, GuildID)
	if err != nil {
		log.Fatal(err)
	}
}
