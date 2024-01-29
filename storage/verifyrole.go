package storage

import "log"

func (db Storage) GetVerifyRole(GuildID string) string {
	var verifyrole string

	// Temporary result until database is implemented
	switch GuildID {
	case ANIMENORTH:
		verifyrole = ATTENDEE
	case TESTSERVER:
		verifyrole = TESTVERIFY
	default:
		verifyrole = ""
	}

	// if err := db.Stmts[verifyrole].QueryRow(GuildID).Scan(&verifyrole); err != nil {
	// 	log.Fatal(err)
	// }

	return verifyrole
}

func (db Storage) __todo__VerifyRole(GuildID string, verifyrole string) {
	_, err := db.Backend.Exec("UPDATE guilds SET verifyrole = ? WHERE guildid = ?", verifyrole, GuildID)
	if err != nil {
		log.Fatal(err)
	}
}
