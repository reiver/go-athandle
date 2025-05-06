package athandle

import (
	"net/http"

	"github.com/reiver/go-erorr"
)

const (
	errNilHTTPRequest       = erorr.Error("athandle: nil http-request")
	errNilHTTPRequestHeader = erorr.Error("athandle: nil http-request header")
)

// useragent holds the default HTTP User-Agent included in HTTP requests.
//
// Its initial value is what is set below.
//
// Its value can be changed by calling athandle.SetUserAgent()
var useragent string = "go-athandle/0.0 (+https://github.com/reiver/go-athandle)"

// SetUserAgent sets the User-Agent in any HTTP requests,
// such as the HTTP request for the DID well-known URL.
func SetUserAgent(value string) {
	useragent = value
}

func setUserAgent(req *http.Request, useragent string) error {
	if nil == req {
		return errNilHTTPRequest
	}

	if "" == useragent {
		return nil
	}

	var header http.Header = req.Header
	{
		if nil == header {
			return errNilHTTPRequestHeader
		}
	}

	header.Add("User-Agent", useragent)
	return nil
}
