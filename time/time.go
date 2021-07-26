package time

import (
	t "time"
)

func NowGit() string {
	return t.Now().Format("Mon Jan 2 15:04:05 MST 2006")
}

func NowUnixMillis() int64 {
	return t.Now().UnixNano() / 1000000
}

func ParseMs(ms int64) t.Time {
	s := ms / 1000
	rem := ms - (s * 1000)
	return t.Unix(s, rem)
}
