package xboxapi

import (
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var DebugHTTP = false

type rt struct {
	cfg    *Config
	client *http.Client
}

func (rt *rt) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("X-Auth", rt.cfg.APIKey)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept-Language", rt.cfg.Language)
	if DebugHTTP {
		reqBuf, _ := httputil.DumpRequestOut(r, true)
		os.Stderr.WriteString("Request:\n--------\n\n")
		os.Stderr.Write(reqBuf)
		start := time.Now()
		rsp, err := rt.client.Do(r)
		took := time.Now().Sub(start).String()
		if err != nil {
			os.Stderr.WriteString("Response:\n--------\n\nError: " + err.Error())
		} else {
			rspBuf, _ := httputil.DumpResponse(rsp, true)
			os.Stderr.WriteString("Response:\n--------\n\n")
			os.Stderr.Write(rspBuf)
		}
		os.Stderr.WriteString("\n\nResponse Time: " + took + "\n\n")
		return rsp, err
	}
	return http.DefaultClient.Do(r)
}
