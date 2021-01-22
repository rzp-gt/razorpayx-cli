package commands

import (
	"bytes"
	"fmt"
	"github.com/rzp-gt/razorpayx-cli/internal/ansi"
	"github.com/rzp-gt/razorpayx-cli/internal/validators"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// create a handler struct
type HttpHandler struct{}

// global variable
var forward_to string

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// create response binary data
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		panic(err)
	}

	// write `data` to response
	result := ansi.ColorizeJSON(string(data), false, os.Stdout)
	fmt.Println(result)

	// Because go lang is a pain in the ass if you read the body then any susequent calls
	// are unable to read the body again....
	req.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	if forward_to != "" {
		serveReverseProxy(forward_to, res, req)
	}

	if err != nil {
		msg := "error"
		res.Write([]byte(msg))
		res.WriteHeader(http.StatusInternalServerError)
	}

	msg := "success"
	res.Write([]byte(msg))

}

type listenCmd struct {
	cmd             *cobra.Command
	forwardURL      string
	livemode        bool
	skipVerify      bool
	apiBaseURL      string
	printJSON       bool
	onlyPrintSecret bool
}

func newListenCmd() *listenCmd {
	lc := &listenCmd{}

	msg := "The listen command watches webhook events from RazorpayX API to your local machine\n" +
		" by connecting directly to RazorpayX's API"

	lc.cmd = &cobra.Command{
		Use:     "listen",
		Args:    validators.NoArgs,
		Short:   ansi.ColoredBoldStatus("Listen for webhook events"),
		Long:    ansi.ColoredBoldStatus(msg),
		Example: `RazorpayX listen`,
		RunE:    lc.runListenCmd,
	}

	lc.cmd.Flags().BoolVar(&lc.livemode, "live", false, "Receive live events (default: test)")
	lc.cmd.Flags().StringVarP(&lc.forwardURL, "forward-to", "f", "", "The URL to forward webhook events to")
	return lc
}

func (lc *listenCmd) runListenCmd(cmd *cobra.Command, args []string) error {

	fmt.Printf("Listening Now \n")

	forward_to = lc.forwardURL

	// create a new handler
	handler := HttpHandler{}
	http.HandleFunc("/webhooks", handler.ServeHTTP)

	// listen and serve
	http.ListenAndServe(":8080", handler)
	return nil
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}
