package xboxapi

import (
	"fmt"
	"net/http"
)

var ErrNotFound = fmt.Errorf("404 Not Found")

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
	switch rsp.StatusCode {
	case 200:
		return nil
	case 404:
		return ErrNotFound
	}
	return httpError{rsp.StatusCode, rsp.Status}
}
