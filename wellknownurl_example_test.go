package athandle_test

import (
	"fmt"

	"github.com/reiver/go-athandle"
)

func ExampleWellKnownURL() {

	var handle string = "Reiver.BSky.Social"

	var url string = athandle.WellKnownURL(handle)

	fmt.Printf("handle:      %s\n", handle)
	fmt.Printf("url: %s\n", url)

	// Output:
	// handle:      Reiver.BSky.Social
	// url: https://reiver.bsky.social/.well-known/atproto-did
}
