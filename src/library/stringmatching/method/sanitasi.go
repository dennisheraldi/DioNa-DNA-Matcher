package method

import (
	"regexp"
)

func sanitasi(p string) bool {
	sanitasi, _ := regexp.Compile("[^AGTC]")
	matchResult := sanitasi.FindString(p)
	if len(matchResult) != 0 {
		return false
	} else {
		return true
	}
}
