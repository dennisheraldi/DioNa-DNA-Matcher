package library

import (
	"regexp"
	"strings"
)

func Sanitasi(p string) bool {
	sanitasi, _ := regexp.Compile("[^AGTC]")
	if sanitasi.MatchString(p) {
		return false
	} else {
		return true
	}
}

func QueryCheck(p string) (string, string) {
	reF, _ := regexp.Compile(`^\d{4}-\d{2}-\d{2}$`)
	reC,_ := regexp.Compile(`^\d{4}-\d{2}-\d{2}\s`)
	reAll,_ := regexp.Compile(`^[aA][lL][lL]$`)

	if (len(p)== 0) {
		return "",""
	} else if (reAll.MatchString(p)) {
		return "all", "all"
	} else if (reF.MatchString(p)) {
		return p, ""
	} else if (reC.MatchString(p)) {
		date := " "
		penyakit := " "
		sepID := strings.Index(p, " ")
		date = p[:sepID]
		penyakit = p[sepID+1:]
		return date, penyakit
	} else {
		return "", p
	}
}
