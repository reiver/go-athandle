package athandle_test

import (
	"testing"

	"github.com/reiver/go-athandle"
)

func TestValidate(t *testing.T) {

	tests := []struct{
		Handle string
		ExpectedError string
	}{
		{
			Handle:                          "",
			ExpectedError: `athandle: handle "" is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                          "once",
			ExpectedError: `athandle: handle "once" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "twice",
			ExpectedError: `athandle: handle "twice" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "thrice",
			ExpectedError: `athandle: handle "thrice" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "fource",
			ExpectedError: `athandle: handle "fource" is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                          "ONCE",
			ExpectedError: `athandle: handle "ONCE" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "Twice",
			ExpectedError: `athandle: handle "Twice" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "tHrICe",
			ExpectedError: `athandle: handle "tHrICe" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "fourcE",
			ExpectedError: `athandle: handle "fourcE" is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                          "host.alt",
			ExpectedError: `athandle: handle "host.alt" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.arpa",
			ExpectedError: `athandle: handle "host.arpa" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.example",
			ExpectedError: `athandle: handle "host.example" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.internal",
			ExpectedError: `athandle: handle "host.internal" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.invalid",
			ExpectedError: `athandle: handle "host.invalid" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.local",
			ExpectedError: `athandle: handle "host.local" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.localhost",
			ExpectedError: `athandle: handle "host.localhost" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.onion",
			ExpectedError: `athandle: handle "host.onion" is invalid because it contains a disallowed TLD`,
		},



		{
			Handle:                          "host.ALT",
			ExpectedError: `athandle: handle "host.ALT" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.ARPA",
			ExpectedError: `athandle: handle "host.ARPA" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.EXAMPLE",
			ExpectedError: `athandle: handle "host.EXAMPLE" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.INTERNAL",
			ExpectedError: `athandle: handle "host.INTERNAL" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.INVALID",
			ExpectedError: `athandle: handle "host.INVALID" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.LOCAL",
			ExpectedError: `athandle: handle "host.LOCAL" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.LOCALHOST",
			ExpectedError: `athandle: handle "host.LOCALHOST" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "host.ONION",
			ExpectedError: `athandle: handle "host.ONION" is invalid because it contains a disallowed TLD`,
		},



		{
			Handle:                          "jo@hn.test",
			ExpectedError: `athandle: handle "jo@hn.test" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "💩.test",
			ExpectedError: `athandle: handle "💩.test" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "john..test",
			ExpectedError: `athandle: handle "john..test" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "xn--bcher-.tld",
			ExpectedError: `athandle: handle "xn--bcher-.tld" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "john.0",
			ExpectedError: `athandle: handle "john.0" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "cn.8",
			ExpectedError: `athandle: handle "cn.8" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "www.masełkowski.pl.com",
			ExpectedError: `athandle: handle "www.masełkowski.pl.com" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "org",
			ExpectedError: `athandle: handle "org" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "name.org.",
			ExpectedError: `athandle: handle "name.org." is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                          "jo@hn.TEST",
			ExpectedError: `athandle: handle "jo@hn.TEST" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "💩.TEST",
			ExpectedError: `athandle: handle "💩.TEST" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "john..TEST",
			ExpectedError: `athandle: handle "john..TEST" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                          "xn--bcher-.TLD",
			ExpectedError: `athandle: handle "xn--bcher-.TLD" is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                          "2gzyxa5ihm7nsggfxnu52rck2vv4rvmdlkiu3zzui5du4xyclen53wid.onion",
			ExpectedError: `athandle: handle "2gzyxa5ihm7nsggfxnu52rck2vv4rvmdlkiu3zzui5du4xyclen53wid.onion" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "laptop.local",
			ExpectedError: `athandle: handle "laptop.local" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "blah.arpa",
			ExpectedError: `athandle: handle "blah.arpa" is invalid because it contains a disallowed TLD`,
		},



		{
			Handle:                          "2gzyxa5ihm7nsggfxnu52rck2vv4rvmdlkiu3zzui5du4xyclen53wid.ONION",
			ExpectedError: `athandle: handle "2gzyxa5ihm7nsggfxnu52rck2vv4rvmdlkiu3zzui5du4xyclen53wid.ONION" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "laptop.LOCAL",
			ExpectedError: `athandle: handle "laptop.LOCAL" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                          "blah.ARPA",
			ExpectedError: `athandle: handle "blah.ARPA" is invalid because it contains a disallowed TLD`,
		},
	}

	for testNumber, test := range tests {

		err := athandle.Validate(test.Handle)
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
			Handle:                          "example.com",
		},
		{
			Handle:                          "Example.com",
		},
		{
			Handle:                          "example.COM",
		},
		{
			Handle:                          "EXamPle.COM",
		},
		{
			Handle:                          "EXAMPLE.com",
		},



		{
			Handle:                          "xn--939a.kr",
		},
		{
			Handle:                          "xn--939a.Kr",
		},
		{
			Handle:                          "xn--939a.kR",
		},
		{
			Handle:                          "xn--939a.KR",
		},
		{
			Handle:                          "xn--939A.KR",
		},
	}

	for testNumber, test := range tests {

		err := athandle.Validate(test.Handle)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("HANDLE: %s", test.Handle)
			continue
		}
	}
}

