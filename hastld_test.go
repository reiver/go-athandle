package athndl

import (
	"testing"
)

func TestHasTLD(t *testing.T) {

	tests := []struct{
		Handle string
		TLD string
		Expected bool
	}{
		{
			Handle: "example.com",
			TLD:           ".com",
			Expected: true,
		},
		{
			Handle: "example.coM",
			TLD:           ".com",
			Expected: true,
		},
		{
			Handle: "example.cOm",
			TLD:           ".com",
			Expected: true,
		},
		{
			Handle: "example.cOM",
			TLD:           ".com",
			Expected: true,
		},
		{
			Handle: "example.Com",
			TLD:           ".com",
			Expected: true,
		},
		{
			Handle: "example.CoM",
			TLD:           ".com",
			Expected: true,
		},
		{
			Handle: "example.COm",
			TLD:           ".com",
			Expected: true,
		},
		{
			Handle: "example.COM",
			TLD:           ".com",
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := hasTLD(test.Handle, test.TLD)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result it not what was expected." , testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("HANDLE: %q", test.Handle)
			t.Logf("TLD: %q", test.TLD)
			continue
		}
	}
}
