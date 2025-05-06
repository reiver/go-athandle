package athandle_test

import (
	"testing"

	"github.com/reiver/go-athandle"
)

func TestWellKnownURL(t *testing.T) {

	tests := []struct{
		Handle string
		Expected string
	}{
		{
			Handle:           "example.com",
			Expected: "https://example.com/.well-known/atproto-did",
		},
		{
			Handle:           "Example.COM",
			Expected: "https://example.com/.well-known/atproto-did",
		},



		{
			Handle:                 "강.kr",
			Expected: "https://xn--939a.kr/.well-known/atproto-did",
		},
		{
			Handle:                 "강.kR",
			Expected: "https://xn--939a.kr/.well-known/atproto-did",
		},
		{
			Handle:                 "강.Kr",
			Expected: "https://xn--939a.kr/.well-known/atproto-did",
		},
		{
			Handle:                 "강.KR",
			Expected: "https://xn--939a.kr/.well-known/atproto-did",
		},
	}

	for testNumber, test := range tests {

		actual := athandle.WellKnownURL(test.Handle)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual HTTP well-known URL is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("HANDLE: %s", test.Handle)
			continue
		}
	}
}
