package storage

func CloseDB() {
	botCache.connection.Close()
}
