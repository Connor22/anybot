package conf

func (config *AnyGuild) GetJoinRole() string {
	var joinrole string

	// Switch cases and global vars used until bot reaches a point where
	// the bot needs more permanent storage
	switch config.ID {
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
