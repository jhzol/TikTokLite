package util

import (
	"time"
)

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func GetCurrentTimeForString() string {
	currentTime := time.Now()
	return currentTime.Format("200601021504")
}
