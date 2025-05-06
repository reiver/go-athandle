package athandle

import (
	"regexp"

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

// Validate does not [Normalize] the handle before applying the validation rules.
func Validate(handle string) error {
	if !regex.MatchString(handle) {
		return erorr.Errorf("atproto: handle %q is invalid because does not fit the valid atproto-handle syntax", handle)
	}

	for _, tld := range disallowedTLDs {
		if hasTLD(handle, tld) {
			return erorr.Errorf("atproto: handle %q is invalid because it contains a disallowed TLD", handle)
		}
	}

	return nil
}
