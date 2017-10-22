package xboxapi

import (
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var DebugHTTP = false

type rt struct {
	cfg *Config
}

func (rt *rt) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("X-Auth", rt.cfg.APIKey)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept-Language", rt.cfg.Language)
	if DebugHTTP {
		reqBuf, _ := httputil.DumpRequestOut(r, true)
		start := time.Now()
		rsp, err := http.DefaultClient.Do(r)
		took := time.Now().Sub(start).String()
		rspBuf, _ := httputil.DumpResponse(rsp, true)
		os.Stderr.WriteString("Request:\n--------\n\n")
		os.Stderr.Write(reqBuf)
		os.Stderr.WriteString("Response:\n--------\n\n")
		os.Stderr.Write(rspBuf)
		os.Stderr.WriteString("\n\nResponse Time: " + took + "\n\n")
		return rsp, err
	}
	return http.DefaultClient.Do(r)
}
