package goshippo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/orijtech/otils"
)

const (
	baseURL = "https://api.goshippo.com"

	envGoShippoToken = "GOSHIPPO_TOKEN"
)

type Client struct {
	mu sync.RWMutex

	rt http.RoundTripper

	__apiKey string
}

func NewClient(tokens ...string) (*Client, error) {
	token := otils.FirstNonEmptyString(tokens...)
	if token != "" {
		return &Client{__apiKey: token}, nil
	}
	return NewClientFromEnv()
}

var (
	errBlankShippoToken = fmt.Errorf("did not find %q in your environment", envGoShippoToken)
)

func NewClientFromEnv() (*Client, error) {
	token := strings.TrimSpace(os.Getenv(envGoShippoToken))
	if token == "" {
		return nil, errBlankShippoToken
	}
	return &Client{__apiKey: token}, nil
}

func (c *Client) SetAPIKey(key string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	c.__apiKey = key
}

func (c *Client) SetHTTPRoundTripper(rt http.RoundTripper) {
	c.mu.Lock()
	c.rt = rt
	c.mu.Unlock()
}

func (c *Client) httpClient() *http.Client {
	c.mu.RLock()
	rt := c.rt
	c.mu.RUnlock()

	if rt == nil {
		rt = http.DefaultTransport
	}

	return &http.Client{Transport: rt}
}

func (c *Client) apiKey() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.__apiKey
}

func (c *Client) doAuthAndReq(req *http.Request) ([]byte, http.Header, error) {
	req.Header.Set("Authorization", fmt.Sprintf("ShippoToken %s", c.apiKey()))
	res, err := c.httpClient().Do(req)
	if err != nil {
		return nil, nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	if !otils.StatusOK(res.StatusCode) {
		return nil, res.Header, errors.New(res.Status)
	}
	slurp, err := ioutil.ReadAll(res.Body)
	return slurp, res.Header, err
}
