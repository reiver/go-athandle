package athandle

import (
	"fmt"
	"net"
)

// lookupTXTError wraps any error we get back from net.LookupTXT().
func lookupTXTError(err error, txtdomainname string) error {
	if nil == err {
		return nil
	}

	switch casted := err.(type) {
	case *net.DNSError:
		switch {
		case casted.IsNotFound:
			return DIDTXTNotFoundError{
				txtdomainname:txtdomainname,
				err:err,
			}
		}
	}

	return fmt.Errorf("athandle: problem looking-up DNS TXT record(s) for domnain-name %q: %w", txtdomainname, err)
}
