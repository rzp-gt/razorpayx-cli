package commands

import (
	"fmt"
	"github.com/rzp-gt/razorpayx-cli/internal/ansi"
	"github.com/rzp-gt/razorpayx-cli/internal/validators"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// create response binary data
	data, err := ioutil.ReadAll(req.Body)

	// write `data` to response
	result := ansi.ColorizeJSON(string(data), false, os.Stdout)
	fmt.Println(result)

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
		Use:   "listen",
		Args:  validators.NoArgs,
		Short: "Listen for webhook events",
		Long: ansi.ColoredBoldStatus(msg),
		Example: `RazorpayX listen`,
		RunE:    lc.runListenCmd,
	}

	lc.cmd.Flags().BoolVar(&lc.livemode, "live", false, "Receive live events (default: test)")
	return lc
}

func (lc *listenCmd) runListenCmd(cmd *cobra.Command, args []string) error {

	fmt.Printf("Listening Now \n")
	
	// create a new handler
	handler := HttpHandler{}

	http.HandleFunc("/webhooks", handler.ServeHTTP)
	// listen and serve
	http.ListenAndServe(":8080", handler)
	return nil
}
