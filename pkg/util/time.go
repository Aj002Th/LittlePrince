package util

import "time"

func FmtTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ParseTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

func NeedToWait(repeat int) time.Duration {
	if repeat <= 0 {
		return 0
	}

	return time.Hour * time.Duration(24/repeat)
}
