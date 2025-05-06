package athandle

// TXTSubDomain returns the sub-domain that would (potentially) store the DID.
//
// With the AT-Protocol, an Internet-domain (such as "example.com") can be used as a 'handle'.
// However, to do this, the Internet-domain (such as "example.com") needs to resolve to a DID.
//
// There are currently 2 separate methods to resolve an Internet-domain (such as "example.com") to a DID.
//
// One of those methods involved a DNS TXT record.
//
// With the AT-Protocol, the DID is stored in a DNS TXT record on a particular sub-domain of the Internet-domain.
//
// If, for example, the Internet-domain is "example.com", then the sub-domain is "_atproto.example.com".
// And if, for example, the Internet-domain is "reiver.link", then the sub-domain is "_atproto.reiver.link".
// And if, for example, the Internet-domain is "once.twice.thrice.fource.xyz", then the sub-domain is "_atproto.once.twice.thrice.fource.xyz".
//
// Here we construct the sub-domain that will be used in the DNS TXT record lookup.
//
// [TXTSubDomain] is called internally by [TXTLookup].
// You should NOT call pass the sub-domain returned by TXTSubDomain to [TXTLookup].
func TXTSubDomain(handle string) string {
	return "_atproto." + handle
}
