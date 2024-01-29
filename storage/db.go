package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// AnimeNorth Server values
const (
	ANIMENORTH = "734513822492131438"
	UNVERIFIED = "1181029814333865994"
	ATTENDEE   = "734514238688854130"
)

// Test server values
const (
	TESTSERVER = "1174354483799654510"
	TESTJOIN   = "1201556569163305020"
	TESTVERIFY = "1201559943816425534"
)

// Values for a future hashmap of prepared SQL statements
// const (
// 	JOINROLE = 1
// )
// var statements = map[int]string{
// 	JOINROLE: "select joinrole from guilds where id = ?",
// }

type Storage struct {
	Backend *sql.DB
	Stmts   map[int]*sql.Stmt
}

func InitDB() *Storage {
	db := new(Storage)
	db.Stmts = make(map[int]*sql.Stmt)
	connection, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}

	// for key, query := range statements {
	// 	log.Println(key)
	// 	log.Println(query)
	// 	db.Stmts[key], err = connection.Prepare(query)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	db.Backend = connection

	return db
}
