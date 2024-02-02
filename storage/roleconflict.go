package storage

import "log"

func (db Cache) __todo__VerifyRole(GuildID string, verifyrole string) {
	_, err := db.connection.Exec("UPDATE guilds SET verifyrole = ? WHERE guildid = ?", verifyrole, GuildID)
	if err != nil {
		log.Fatal(err)
	}
}
