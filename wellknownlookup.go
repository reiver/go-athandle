package athandle

import (
	gobytes "bytes"
	"io"
	"net/http"

	"github.com/reiver/go-erorr"
)

// WellKnownLookup returns any DIDs registered for the domain-name, using the HTTP well-known method.
//
// For example:
//
//	did, err := athandle.WellKnownLookup("example.com")
//
// If you are not sure whether to use WellKnownLookup or [Resolve], use [Resolve].
func WellKnownLookup(handle string) (string, error) {

	var wellknownurl string = WellKnownURL(handle)

	var req *http.Request
	{
		var err error
		req, err = http.NewRequest(http.MethodGet, wellknownurl, nil)
		if nil != err {
			return "", erorr.Errorf("athandle: problem creating HTTP-request for HTTP well-known URL %q for handle %q: %w", wellknownurl, handle, err)
		}

		req.Header.Add("Accept", "text/plain")
		req.Header.Add("Connection", "close")
//@TODO: make the User-Agent configurable.
		req.Header.Add("User-Agent", "github.com/reiver/go-athandle")
	}


	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		return "", erorr.Errorf("athandle: problem looking up DID using HTTP well-known URL %q for handle %q: %w", wellknownurl, handle, err)
	}
	if nil == resp {
		return "", erorr.Errorf("athandle: problem looking up DID using HTTP well-known URL %q for handle %q: nil HTTP response", wellknownurl, handle)
	}
	if http.StatusOK != resp.StatusCode {
		return "", erorr.Errorf("athandle: problem looking up DID using HTTP well-known URL %q for handle %q: HTTP response status-code %d", wellknownurl, handle, resp.StatusCode)
	}
	if nil == resp.Body {
		return "", erorr.Errorf("athandle: problem looking up DID using HTTP well-known URL %q for handle %q: nil HTTP response body", wellknownurl, handle)
	}

	var did string
	{
		bytes, err := io.ReadAll(resp.Body)
		if nil != err {
			return "", erorr.Errorf("athandle: problem reading HTTP response body (which should contain the DID) of HTTP well-known URL %q for handle %q: %w", wellknownurl, handle, err)
		}

		did = string(gobytes.TrimSpace(bytes))
	}

	return did, nil
}
