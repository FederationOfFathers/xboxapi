package xboxapi

import (
	"fmt"
	"net/http"
)

type httpError struct {
	StatusCode int
	Status     string
}

func isHTTPError(err error) bool {
	switch err.(type) {
	case httpError:
		return true
	}
	return false
}

func (h httpError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", h.StatusCode, h.Status)
}

func rspError(rsp *http.Response) error {
	if rsp.StatusCode == 200 {
		return nil
	}
	return httpError{rsp.StatusCode, rsp.Status}
}
