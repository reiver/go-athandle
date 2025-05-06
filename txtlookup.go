package athandle

import (
	"fmt"
	"net"
)

// TXTLookup returns any DIDs registered for the domain-name, using the special DNS TXT method.
//
// Note that, internally TXTLookup calls [TXTSubDomain].
// You should not call [TXTSubDomain] yourself, if you are calling TXTLookup.
//
// For example:
//
//	dids, err := athandle.TXTLookup("example.com")
//
// If you are not sure whether to use TXTLookup or [Resolve], use [Resolve].
func TXTLookup(handle string) ([]string, error) {

	var txtdomainname string = TXTSubDomain(handle)

	// Return all the DNS TXT records for the "_atproto" sub-domain.
	//
	// Note that these could contain records that are NOT DIDs.
	// For example:
	//
	//	did=did:plc:m2jwplpernhxkzbo4ev5ljwj
	//	once=apple
	//	twice=banana
	//	thrice=cherry
	//	fource=date
	//	did=did:something:abcde12345
	//
	// Later we will filter out the non-DIDs.
	var txtresponses []string
	{
		var err error

		txtresponses, err = net.LookupTXT(txtdomainname)
		if nil != err {
			err = lookupTXTError(err, txtdomainname)
			return nil, err
		}
	}

	var dids []string = filterDIDs(txtresponses)
	if len(dids) < 1 {
		return nil, DIDTXTNotFoundError{
			txtdomainname:txtdomainname,
			err:fmt.Errorf("athandle: no \"did=\" DNS TXT records for %q", txtdomainname),
		}
	}

	return dids, nil
}
