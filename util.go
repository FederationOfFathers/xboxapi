package xboxapi

import "net/http"

func safeHTTPResponseClose(r *http.Response) {
	if r == nil {
		return
	}
	if r.Body == nil {
		return
	}
	r.Body.Close()
}
