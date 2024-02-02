package conf

func (config *AnyGuild) GetVerifyRole() string {
	var verifyrole string

	// Temporary result until database is implemented
	switch config.ID {
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
