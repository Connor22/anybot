package storage

import "log"

func (db Storage) GetJoinRole(GuildID string) string {
	var joinrole string

	// Switch cases and global vars used until bot reaches a point where
	// the bot needs more permanent storage
	switch GuildID {
	case ANIMENORTH:
		joinrole = UNVERIFIED
	case TESTSERVER:
		joinrole = TESTJOIN
	default:
		joinrole = ""
	}

	// if err := db.Stmts[JOINROLE].QueryRow(GuildID).Scan(&joinrole); err != nil {
	// 	log.Fatal(err)
	// }

	return joinrole
}

func (db Storage) __todo__SetJoinRole(GuildID string, joinrole string) {
	_, err := db.Backend.Exec("UPDATE guilds SET joinrole = ? WHERE guildid = ?", joinrole, GuildID)
	if err != nil {
		log.Fatal(err)
	}
}
