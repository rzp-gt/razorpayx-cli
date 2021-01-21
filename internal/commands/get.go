package commands

import (
	"github.com/rzp-gt/razorpayx-cli/internal/requests"
	"github.com/rzp-gt/razorpayx-cli/internal/validators"
	"github.com/spf13/cobra"
	"net/http"
)

type getCmd struct {
	reqs requests.Base
}

func newGetCmd() *getCmd {
	gc := &getCmd{}

	gc.reqs.Method = http.MethodGet
	gc.reqs.Profile = &Config.Profile
	gc.reqs.Cmd = &cobra.Command{
		Use:   "get <id or path>",
		Args:  validators.ExactArgs(1),
		Short: "Retrieve resources by their ID or make GET requests",
		Long: `With the get command, you can load API resources by providing just the resource
			   id. You can also make normal HTTP GET requests to the Stripe API by providing
 			   the API path.`,
		Example: `razorpay get pout_GLBIjRm3dN3i4Y`,
		RunE: gc.reqs.RunRequestsCmd,
	}

	gc.reqs.InitFlags()

	return gc
}
