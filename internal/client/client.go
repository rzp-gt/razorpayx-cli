package client

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// DefaultAPIBaseURL is the default base URL for API requests
const DefaultAPIBaseURL = "https://api.razorpay.com/"

// DefaultDashboardBaseURL is the default base URL for dashboard requests
const DefaultDashboardBaseURL = "https://x.razorpay.com/"

// Client is the API client used to sent requests to Stripe.
type Client struct {
	// The base URL (protocol + hostname) used for all requests sent by this
	// client.
	BaseURL *url.URL

	// API key used to authenticate requests sent by this client. If left
	// empty, the `Authorization` header will be omitted.
	APIKey string

	APISecret string

	// When this is enabled, request and response headers will be printed to
	// stdout.
	Verbose bool

	// Cached HTTP client, lazily created the first time the Client is used to
	// send a request.
	httpClient *http.Client
}

// PerformRequest sends a request to Stripe and returns the response.
func (c *Client) PerformRequest(ctx context.Context, method, path string, params string, configure func(*http.Request)) (*http.Response, error) {
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	url = c.BaseURL.ResolveReference(url)

	var body io.Reader
	if method == http.MethodPost {
		body = strings.NewReader(params)
	} else {
		url.RawQuery = params
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//req.SetBasicAuth("rzp_test_KAtANE6dy27Q5D","bj7PJz2TbOJ90HhgcubsNezm" )

	req.SetBasicAuth(c.APIKey, c.APISecret)

	if configure != nil {
		configure(req)
	}

	if c.httpClient == nil {
		c.httpClient = newHTTPClient(c.Verbose, os.Getenv("RAZORPAY_X_UNIX_SOCKET"))
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func newHTTPClient(verbose bool, unixSocket string) *http.Client {
	var httpTransport *http.Transport

	if unixSocket != "" {
		dialFunc := func(network, addr string) (net.Conn, error) {
			return net.Dial("unix", unixSocket)
		}
		dialContext := func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", unixSocket)
		}
		httpTransport = &http.Transport{
			DialContext:           dialContext,
			DialTLS:               dialFunc,
			ResponseHeaderTimeout: 30 * time.Second,
			ExpectContinueTimeout: 10 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
		}
	} else {
		httpTransport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
		}
	}

	return &http.Client{
		Transport: httpTransport,
	}
}
