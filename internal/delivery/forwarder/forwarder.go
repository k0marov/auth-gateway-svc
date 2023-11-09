package forwarder

import (
	"net/http"
	"net/http/httputil"
)

func NewForwarder(forwardHost string) http.Handler {
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.Host = forwardHost
			r.URL.Host = forwardHost
			r.URL.Scheme = determineScheme(r)
		},
	}
}

// Since strangely req.URL.Scheme is always empty, you need this hack for the proxy to work
// refer to https://github.com/golang/go/issues/28940
// Possible workaround: providing scheme in the proxy config itself
func determineScheme(req *http.Request) string {
	if req.TLS != nil {
		return "https"
	}
	return "http"
}
