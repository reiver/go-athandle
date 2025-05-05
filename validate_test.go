package athndl_test

import (
	"testing"

	"github.com/reiver/go-athndl"
)

func TestValidate(t *testing.T) {

	tests := []struct{
		Handle string
		ExpectedError string
	}{
		{
			Handle:                         "",
			ExpectedError: `atproto: handle "" is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                         "once",
			ExpectedError: `atproto: handle "once" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "twice",
			ExpectedError: `atproto: handle "twice" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "thrice",
			ExpectedError: `atproto: handle "thrice" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "fource",
			ExpectedError: `atproto: handle "fource" is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                         "ONCE",
			ExpectedError: `atproto: handle "ONCE" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "Twice",
			ExpectedError: `atproto: handle "Twice" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "tHrICe",
			ExpectedError: `atproto: handle "tHrICe" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "fourcE",
			ExpectedError: `atproto: handle "fourcE" is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                         "host.alt",
			ExpectedError: `atproto: handle "host.alt" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.arpa",
			ExpectedError: `atproto: handle "host.arpa" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.example",
			ExpectedError: `atproto: handle "host.example" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.internal",
			ExpectedError: `atproto: handle "host.internal" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.invalid",
			ExpectedError: `atproto: handle "host.invalid" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.local",
			ExpectedError: `atproto: handle "host.local" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.localhost",
			ExpectedError: `atproto: handle "host.localhost" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.onion",
			ExpectedError: `atproto: handle "host.onion" is invalid because it contains a disallowed TLD`,
		},
	}

	for testNumber, test := range tests {

		err := athndl.Validate(test.Handle)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("HANDLE: %s", test.Handle)
			continue
		}

		actual := err.Error()
		expected := test.ExpectedError

		if expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected." , testNumber)
			t.Logf("EXPECTED-ERROR: %q", expected)
			t.Logf("ACTUAL-ERROR:   %q", actual)
			t.Logf("HANDLE: %q", test.Handle)
			continue
		}
	}
}

func TestValidate_success(t *testing.T) {

	tests := []struct{
		Handle string
	}{
		{
			Handle:                         "example.com",
		},
		{
			Handle:                         "Example.com",
		},
		{
			Handle:                         "example.COM",
		},
		{
			Handle:                         "EXamPle.COM",
		},
		{
			Handle:                         "EXAMPLE.com",
		},



		{
			Handle:                         "강.kr",
		},
		{
			Handle:                         "강.Kr",
		},
		{
			Handle:                         "강.kR",
		},
		{
			Handle:                         "강.KR",
		},
		{
			Handle:                         "xn--939A.KR",
		},
	}

	for testNumber, test := range tests {

		err := athndl.Validate(test.Handle)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("HANDLE: %s", test.Handle)
			continue
		}
	}
}

