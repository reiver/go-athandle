package athandle_test

import (
	"fmt"

	"github.com/reiver/go-athandle"
)

func ExampleTXTSubDomain() {

	var domain string = "example.com"

	subdomain := athandle.TXTSubDomain(domain)

	fmt.Printf("domain:              %s\n", domain)
	fmt.Printf("sub-domain: %s\n", subdomain)

	// Output:
	// domain:              example.com
	// sub-domain: _atproto.example.com
}
