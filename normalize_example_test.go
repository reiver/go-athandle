package athandle_test

import (
	"fmt"

	"github.com/reiver/go-athandle"
)

func ExampleNormalize() {

	var handle string = "Reiver.BSky.Social"

	normalized := athandle.Normalize(handle)

	fmt.Printf("handle:            %s\n", handle)
	fmt.Printf("normalized handle: %s\n", normalized)

	// Output:
	// handle:            Reiver.BSky.Social
	// normalized handle: reiver.bsky.social
}
