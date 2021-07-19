package geziyorhelper

import (
	"strings"

	"github.com/geziyor/geziyor/client"
)

const (
	DisableDefaultHeadersKey = "disable_default_headers"
)

const (
	headerSep    = ": "
	headerSepLen = len(headerSep)
	minLineLen   = 4
)

// DefaultHeaders sets default request headers, it's more flexible.
// if r.Meta[DisableDefaultHeadersKey] exists, it won't set default headers.
type DefaultHeaders struct {
	Headers map[string]string
}

// LoadHeaders loads headers from a string that you can copy from chrome network
func LoadHeaders(s string) map[string]string {
	headers := make(map[string]string)
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		if len(line) < minLineLen {
			continue
		}
		pos := strings.Index(line, headerSep)
		if pos <= 0 {
			continue
		}

		headers[line[0:pos]] = line[pos+headerSepLen:]
	}

	return headers
}

func (d *DefaultHeaders) ProcessRequest(r *client.Request) {
	if _, ok := r.Meta[DisableDefaultHeadersKey]; !ok {
		for key, value := range d.Headers {
			r.Header = client.SetDefaultHeader(r.Header, key, value)
		}
	}
}
