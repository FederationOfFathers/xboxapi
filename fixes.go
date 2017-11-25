package xboxapi

import (
	"fmt"
	"strings"
	"time"
)

// APITime -- because time is hard and WTF am I supposed to do with "0001-01-01T00:00:00" ?
type APITime string

// Time returns a time.Time or the time of the Unix epoch in case of any parsing error
func (a APITime) Time() time.Time {
	if a == "0001-01-01T00:00:00" {
		return time.Unix(1, 0)
	}
	if a[10] == ' ' {
		// "2012-11-04 06:11:01" -> "2012-11-04T06:11:01.0000000Z"
		a = APITime(fmt.Sprintf("%s.0000000Z", strings.Replace(string(a), " ", "T", 1)))
	}
	t, err := time.Parse(time.RFC3339Nano, string(a))
	if err != nil {
		return time.Unix(1, 0)
	}
	return t
}
