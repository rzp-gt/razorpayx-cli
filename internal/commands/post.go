package commands

import (
	"github.com/rzp-gt/razorpayx-cli/internal/ansi"
	"github.com/rzp-gt/razorpayx-cli/internal/requests"
	"github.com/rzp-gt/razorpayx-cli/internal/validators"
	"github.com/spf13/cobra"
	"net/http"
)

type postCmd struct {
	reqs requests.Base
}

func newPostCmd() *postCmd {
	gc := &postCmd{}

	msg := "Make POST requests to the RazorpayX API using your key. \n" +
		"The post command supports API features like idempotency keys.\n"

	gc.reqs.Method = http.MethodPost
	gc.reqs.Cmd = &cobra.Command{
		Use:   "post <path>",
		Args:  validators.ExactArgs(1),
		RunE:  gc.reqs.RunRequestsCmd,
		Short: ansi.ColoredBoldStatus("Make POST requests to RazorpayX API"),
		Long: ansi.ColoredBoldStatus(msg),
		Example: `razorpayx post /payouts \
    			  -d amount=2000 \
				  -d account_number=2323230079767628 \
		`}
	gc.reqs.Profile = &Config.Profile
	gc.reqs.InitFlags()

	return gc
}
