package coinmarketcap

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var coinmarketcapUrl = "https://pro-api.coinmarketcap.com/"

func NewClient(key string) *Client {
	return &Client{
		APIKey:     key,
		BaseURL:    coinmarketcapUrl,
		UserAgent:  "Chrome/67.0.3396.99",
		HTTPClient: http.DefaultClient,
		Debug:      false,
		Logger:     log.New(os.Stderr, "go-coinmarketcap:", log.LstdFlags),
	}
}

func (c *Client) parseRequest(r *Request, opts ...RequestOption) (err error) {
	// set Request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.Validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.Endpoint)
	queryString := r.Query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.Form.Encode()
	header := http.Header{}
	header.Set("X-CMC_PRO_API_KEY", c.APIKey)
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-Form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}

	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, Body: %s", fullURL, bodyString)

	r.FullURL = fullURL
	r.Header = header
	r.Body = body
	return nil
}

func (c *Client) CallAPI(ctx context.Context, r *Request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.Method, r.FullURL, r.Body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Close = true
	req.Header = r.Header
	c.debug("Request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the Body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response Body: %s", string(data))

	if res.StatusCode != 200 {
		hErr := new(HTTPError)
		hErr.StatusCode = res.StatusCode
		hErr.Body = data
		return nil, hErr
	}

	return data, nil
}

type doFunc func(req *http.Request) (*http.Response, error)

// Client define API client
type Client struct {
	APIKey     string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	do         doFunc
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}
