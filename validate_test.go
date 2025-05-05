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



		{
			Handle:                         "host.ALT",
			ExpectedError: `atproto: handle "host.ALT" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.ARPA",
			ExpectedError: `atproto: handle "host.ARPA" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.EXAMPLE",
			ExpectedError: `atproto: handle "host.EXAMPLE" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.INTERNAL",
			ExpectedError: `atproto: handle "host.INTERNAL" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.INVALID",
			ExpectedError: `atproto: handle "host.INVALID" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.LOCAL",
			ExpectedError: `atproto: handle "host.LOCAL" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.LOCALHOST",
			ExpectedError: `atproto: handle "host.LOCALHOST" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "host.ONION",
			ExpectedError: `atproto: handle "host.ONION" is invalid because it contains a disallowed TLD`,
		},



		{
			Handle:                         "jo@hn.test",
			ExpectedError: `atproto: handle "jo@hn.test" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "💩.test",
			ExpectedError: `atproto: handle "💩.test" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "john..test",
			ExpectedError: `atproto: handle "john..test" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "xn--bcher-.tld",
			ExpectedError: `atproto: handle "xn--bcher-.tld" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "john.0",
			ExpectedError: `atproto: handle "john.0" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "cn.8",
			ExpectedError: `atproto: handle "cn.8" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "www.masełkowski.pl.com",
			ExpectedError: `atproto: handle "www.masełkowski.pl.com" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "org",
			ExpectedError: `atproto: handle "org" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "name.org.",
			ExpectedError: `atproto: handle "name.org." is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                         "jo@hn.TEST",
			ExpectedError: `atproto: handle "jo@hn.TEST" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "💩.TEST",
			ExpectedError: `atproto: handle "💩.TEST" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "john..TEST",
			ExpectedError: `atproto: handle "john..TEST" is invalid because does not fit the valid atproto-handle syntax`,
		},
		{
			Handle:                         "xn--bcher-.TLD",
			ExpectedError: `atproto: handle "xn--bcher-.TLD" is invalid because does not fit the valid atproto-handle syntax`,
		},



		{
			Handle:                         "2gzyxa5ihm7nsggfxnu52rck2vv4rvmdlkiu3zzui5du4xyclen53wid.onion",
			ExpectedError: `atproto: handle "2gzyxa5ihm7nsggfxnu52rck2vv4rvmdlkiu3zzui5du4xyclen53wid.onion" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "laptop.local",
			ExpectedError: `atproto: handle "laptop.local" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "blah.arpa",
			ExpectedError: `atproto: handle "blah.arpa" is invalid because it contains a disallowed TLD`,
		},



		{
			Handle:                         "2gzyxa5ihm7nsggfxnu52rck2vv4rvmdlkiu3zzui5du4xyclen53wid.ONION",
			ExpectedError: `atproto: handle "2gzyxa5ihm7nsggfxnu52rck2vv4rvmdlkiu3zzui5du4xyclen53wid.ONION" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "laptop.LOCAL",
			ExpectedError: `atproto: handle "laptop.LOCAL" is invalid because it contains a disallowed TLD`,
		},
		{
			Handle:                         "blah.ARPA",
			ExpectedError: `atproto: handle "blah.ARPA" is invalid because it contains a disallowed TLD`,
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
			Handle:                         "xn--939a.kr",
		},
		{
			Handle:                         "xn--939a.Kr",
		},
		{
			Handle:                         "xn--939a.kR",
		},
		{
			Handle:                         "xn--939a.KR",
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

