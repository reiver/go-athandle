package athandle

import (
	"strings"
)

// filterDIDs is used to pull out the DIDs from the result of a DNS TXT lookup.
func filterDIDs(txtresponses []string) (dids []string) {

	const prefix string = "did="

	for _, resp := range txtresponses {

		if strings.HasPrefix(resp, prefix) {
			var did string = resp[len(prefix):]

			if "" != did {
				dids = append(dids, did)
			}
		}
	}

	return dids
}
