package athndl_test

import (
	"testing"

	"github.com/reiver/go-athndl"
)

func TestNormalize(t *testing.T) {

	tests := []struct{
		Handle string
		Expected string
	}{
		{
			Handle:   "",
			Expected: "",
		},



		{
			Handle:   "once",
			Expected: "once",
		},
		{
			Handle:   "twice",
			Expected: "twice",
		},
		{
			Handle:   "thrice",
			Expected: "thrice",
		},
		{
			Handle:   "fource",
			Expected: "fource",
		},



		{
			Handle:   "ONCE",
			Expected: "once",
		},
		{
			Handle:   "Twice",
			Expected: "twice",
		},
		{
			Handle:   "tHrICe",
			Expected: "thrice",
		},
		{
			Handle:   "fourcE",
			Expected: "fource",
		},



		{
			Handle:   "example.com",
			Expected: "example.com",
		},
		{
			Handle:   "Example.com",
			Expected: "example.com",
		},
		{
			Handle:   "example.COM",
			Expected: "example.com",
		},
		{
			Handle:   "EXamPle.COM",
			Expected: "example.com",
		},
		{
			Handle:   "EXAMPLE.com",
			Expected: "example.com",
		},



		{
			Handle:   "강.kr",
			Expected: "xn--939a.kr",
		},
		{
			Handle:   "강.Kr",
			Expected: "xn--939a.kr",
		},
		{
			Handle:   "강.kR",
			Expected: "xn--939a.kr",
		},
		{
			Handle:   "강.KR",
			Expected: "xn--939a.kr",
		},
		{
			Handle:   "xn--939A.KR",
			Expected: "xn--939a.kr",
		},



		{
			Handle:   "once.TWICE.tHriCe.FoUrCe.example.COM",
			Expected: "once.twice.thrice.fource.example.com",
		},
	}

	for testNumber, test := range tests {

		actual := athndl.Normalize(test.Handle)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual normalized-handle is not what was expected." , testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("HANDLE:   %q", test.Handle)
			continue
		}
	}
}
