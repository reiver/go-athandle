package athandle

import (
	"fmt"
)

var _ error = DIDNotFoundError{}

// DIDNotFoundError is an error that can be returned by [Resolve] if the "_atproto." sub-domain doesn't exist
// or there are no "did=" DNS TXT records for that sub-domain.
type DIDNotFoundError struct {
	txtdomainname string
	err error
}

func (receiver DIDNotFoundError) Error() string {
	return fmt.Sprintf("athandle: DID not found — DNS TXT domain-name=%q: %v", receiver.txtdomainname, receiver.err)
}

func (receiver DIDNotFoundError) Unwrap() error {
	return receiver.err
}
