package coinmarketcap

import (
	"fmt"
	"time"
)

type HTTPError struct {
	StatusCode int
	Body       []byte
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("http error code %d,Body %s", e.StatusCode, string(e.Body))
}

func IsHTTPError(err error) bool {
	_, ok := err.(*HTTPError)
	return ok
}

type Status struct {
	TimeStamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}

func (s *Status) Success() bool {
	return s.ErrorCode == 0
}
