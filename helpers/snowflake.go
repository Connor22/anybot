package helpers

import (
	"time"
)

const (
	discordEpoch   int64 = 1420070400000 // Discord epoch (Jan 1, 2015) in milliseconds
	workerIDShift        = 17
	processIDShift       = 12
	incrementMask        = 0xFFF
)

var (
	workerID  int64 = 1
	processID int64 = 1
	increment int64 = 0
)

type SnowflakeInfo struct {
	Time      time.Time
	WorkerID  int64
	ProcessID int64
	Increment int64
}

func ReverseSnowflake(snowflake int64) int64 {
	timestamp := (snowflake >> 22) + discordEpoch

	return timestamp
}
