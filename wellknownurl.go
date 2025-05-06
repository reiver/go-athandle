package athandle

import (
	liburl "net/url"
)

// WellKnownURL returns the URL that would (potentially) store the DID.
//
// With the AT-Protocol, an Internet-domain (such as "example.com") can be used as a 'handle'.
// However, to do this, the Internet-domain (such as "example.com") needs to resolve to a DID.
//
// There are currently 2 separate methods to resolve an Internet-domain (such as "example.com") to a DID.
//
// One of those methods involved an HTTP well-known URL.
//
// If, for example, the Internet-domain is "example.com", then the HTTP well-known URL is "https://example.com/.well-known/atproto-did".
// And if, for example, the Internet-domain is "reiver.link", then the HTTP well-known URL is "https://reiver.link/.well-known/atproto-did".
// And if, for example, the Internet-domain is "once.twice.thrice.fource.xyz", then the HTTP well-known URL is "https://once.twice.thrice.fource.xyz/.well-known/atproto-did".
//
// Here we construct the URL that will be used in the HTTP-request.
//
// [WellKnownURL] is called internally by [WellKnownLookup].
// You should NOT call pass the URL returned by WellKnownURL to [WellKnownLookup].
func WellKnownURL(handle string) string {

	const path string = "/.well-known/atproto-did"

	var host string = Normalize(handle)

	var url = liburl.URL{
		Scheme: "https",
		Host: host,
		Path: path,
	}

	return url.String()
}
