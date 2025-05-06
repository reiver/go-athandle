package athandle

import (
	"fmt"
)

var _ error = DIDTXTNotFoundError{}

// DIDTXTNotFoundError is an error that can be returned by [Resolve] if the "_atproto." sub-domain doesn't exist
// or there are no "did=" DNS TXT records for that sub-domain.
type DIDTXTNotFoundError struct {
	txtdomainname string
	err error
}

func (receiver DIDTXTNotFoundError) Error() string {
	return fmt.Sprintf("athandle: DID TXT not found — DNS TXT domain-name=%q: %v", receiver.txtdomainname, receiver.err)
}

func (receiver DIDTXTNotFoundError) Unwrap() error {
	return receiver.err
}
