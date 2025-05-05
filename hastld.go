package athndl

import (
	"strings"
)

func hasTLD(handle string, tld string) bool {
	var lenHandle int = len(handle)
	var lenTLD    int = len(tld)

	if lenHandle < lenTLD {
		return false
	}

	var str string = handle[lenHandle-lenTLD:]

	return strings.ToLower(str) == tld
}
