package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	mt "github.com/dustinsprk/tools/time"
)

var TimeZones = []string{
	"America/New_York",
	"America/Chicago",
	"America/Denver",
	"America/Los_Angeles",
}

func main() {
	locs := []*time.Location{}
	now := time.Now()
	local := now.Location()
	found := false
	for _, tz := range TimeZones {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			errorf("error parsing timezone %s %v", tz, err)
			return
		}
		if loc.String() == local.String() {
			found = true
		}
		locs = append(locs, loc)
	}
	fmt.Printf("UTC:\t\t%s\n", mt.FmtTime(now.UTC()))
	localS := mt.FmtTime(now)
	for _, loc := range locs {
		d := mt.FmtTime(now.In(loc))
		if d == localS {
			found = true
		}
		name := strings.Replace(loc.String(), "America/", "", 1)
		name = strings.Replace(name, "_", " ", -1)
		sep := "\t"
		if len(name) < 7 {
			sep = sep + sep
		}
		fmt.Printf("%s:%s%s\n", name, sep, d)
	}
	if !found {
		fmt.Printf("Local:\t\t%s\n", localS)
	}
}

func errorf(s string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stderr, s, args...)
}
