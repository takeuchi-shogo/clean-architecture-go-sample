package utilities

import "time"

func SetCurrentUnixTime() int64 {
	return time.Now().Unix()
}
