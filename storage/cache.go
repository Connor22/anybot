package storage

import (
	"anybot/conf"
	"anybot/modules"
	"database/sql"
	"log"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

type Cache struct {
	connection *sql.DB
	stmts      map[int]*sql.Stmt
	guilds     map[string]*conf.AnyGuild
	Modules    []modules.Module
}

var botCache Cache

func InitCache() *Cache {
	cache := new(Cache)
	cache.stmts = make(map[int]*sql.Stmt)

	connection, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}

	cache.Modules = modules.InitilializeModules()

	cache.connection = connection

	return cache
}

func GetCache() *Cache {
	return &botCache
}

func getGuild(discord *discordgo.Session, gid string) *conf.AnyGuild {
	guild, err := discord.State.Guild(gid)
	if err != nil {
		log.Fatal(err)
	}

	config := &conf.AnyGuild{
		Name: guild.Name,
		ID:   gid,
	}

	config.Flags = 0

	//TODO - fetch config from backend
	switch gid {
	case conf.ANIMENORTH:
		ToggleFlagForMod(config, "JoinRole")
	case conf.TESTSERVER:
		ToggleFlagForMod(config, "JoinRole", "RoleConflict")
	}

	return config
}

func (cache *Cache) GetGuild(discord *discordgo.Session, gid string) *conf.AnyGuild {
	config, found := cache.guilds[gid]

	if !found {
		cache.guilds[gid] = getGuild(discord, gid)
		config = cache.guilds[gid]
	}

	return config
}

func ToggleFlagForMod(config *conf.AnyGuild, toEnable ...string) {
	// POTENTIAL: figure out how to avoid double nested loop
	for _, moduleName := range toEnable {
		for _, module := range botCache.Modules {
			if moduleName == module.Name() {
				config.Flags |= module.Flag()
			}
		}
	}
}
