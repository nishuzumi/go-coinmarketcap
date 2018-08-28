package coinmarketcap

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Params map[string]interface{}

// Request define an API Request
type Request struct {
	Method   string
	Endpoint string
	Query    url.Values
	Form     url.Values
	Header   http.Header
	Body     io.Reader
	FullURL  string
}

// RequestOption define option type for Request
type RequestOption func(*Request)

// setParam set param with key/value to Query string
func (r *Request) SetParam(key string, value interface{}) *Request {
	if r.Query == nil {
		r.Query = url.Values{}
	}
	r.Query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setParams set Params with key/values to Query string
func (r *Request) SetParams(m Params) *Request {
	for k, v := range m {
		r.SetParam(k, v)
	}
	return r
}

// setFormParam set param with key/value to Request Form Body
func (r *Request) SetFormParam(key string, value interface{}) *Request {
	if r.Form == nil {
		r.Form = url.Values{}
	}
	r.Form.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setFormParams set Params with key/values to Request Form Body
func (r *Request) SetFormParams(m Params) *Request {
	for k, v := range m {
		r.SetFormParam(k, v)
	}
	return r
}

func (r *Request) Validate() (err error) {
	if r.Query == nil {
		r.Query = url.Values{}
	}
	if r.Form == nil {
		r.Form = url.Values{}
	}
	return nil
}
