package athndl

import (
	"regexp"
	"strings"

	"github.com/reiver/go-erorr"
)

var disallowedTLDs = []string{
	".alt",
	".arpa",
	".example",
	".internal",
	".invalid",
	".local",
	".localhost",
	".onion",
}

// This regular-expression comes from the AT-Protocol handle specification: https://atproto.com/specs/handle
var regex *regexp.Regexp = regexp.MustCompile(
	`^`+
		`(`+
			`[a-zA-Z0-9]`+
			`(`+
				`[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]`+
			`)?`+
			`\.`+
		`)+`+
		`[a-zA-Z]`+
		`(`+
			`[a-zA-Z0-9-]{0,61}`+
			`[a-zA-Z0-9]`+
		`)?`+
	`$`,
)

func Validate(handle string) error {
	var normalized string = Normalize(handle)

	if !regex.MatchString(normalized) {
		return erorr.Errorf("atproto: handle %q is invalid because does not fit the valid atproto-handle syntax", handle)
	}

	for _, tld := range disallowedTLDs {
		if strings.HasSuffix(normalized, tld) {
			return erorr.Errorf("atproto: handle %q is invalid because it contains a disallowed TLD", handle)
		}
	}

	return nil
}
