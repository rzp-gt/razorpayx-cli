package commands

import (
	"fmt"
	"github.com/rzp-gt/razorpayx-cli/internal/ansi"
	"sort"
	"strings"

	"github.com/rzp-gt/razorpayx-cli/internal/wiki"
	"github.com/spf13/cobra"
)

var nameURLmap = map[string]string{
	"contact":      		"https://razorpay.com/docs/razorpayx/api/contacts",
	"fundAccount":  		"https://razorpay.com/docs/razorpayx/api/fund-accounts",
	"payouts":      		"https://razorpay.com/docs/razorpayx/api/payouts",
	"compositeApi": 		"https://razorpay.com/docs/razorpayx/api/composite-api",
	"idempotency":  		"https://razorpay.com/docs/razorpayx/api/idempotency",
	"transactions": 		"https://razorpay.com/docs/razorpayx/api/transactions",
	"webhooks":     		"https://razorpay.com/docs/razorpayx/api/webhooks",
	"razorpayx":			"https://razorpay.com/docs/razorpayx/",
	"payout-links":			"https://razorpay.com/docs/razorpayx/api/payout-links/",
	"postman-collection":	"https://app.getpostman.com/run-collection/70913b179a2d97235f6f",
	"vendor-payments":		"https://razorpay.com/docs/razorpayx/vendor-payments/",
	"tax-payments":			"https://razorpay.com/docs/razorpayx/tax-payments/",
	"test-mode":			"https://razorpay.com/docs/razorpayx/test-mode/",
	"security":				"https://razorpay.com/docs/razorpayx/2fa/",
}

func wikiNames() []string {
	keys := make([]string, 0, len(nameURLmap))
	for k := range nameURLmap {
		keys = append(keys, k)
	}

	return keys
}

func getLongestShortcut(shortcuts []string) int {
	longest := 0
	for _, s := range shortcuts {
		if len(s) > longest {
			longest = len(s)
		}
	}

	return longest
}

func padName(name string, length int) string {
	difference := length - len(name)

	var b strings.Builder

	fmt.Fprint(&b, name)

	for i := 0; i < difference; i++ {
		fmt.Fprint(&b, " ")
	}

	return b.String()
}

type wikiCmd struct {
	cmd      *cobra.Command
	livemode bool
}

func newWikiCmd() *wikiCmd {
	wc := &wikiCmd{}

	msg := "The wiki command provices shortcuts to quickly let you open pages to RazorpayX with in your browser\n" +
		"A full list of support shortcuts can be seen with 'RazorpayX wiki --list'"

	wc.cmd = &cobra.Command{
		Use:       "wiki",
		ValidArgs: wikiNames(),
		Short:     "Quickly open RazorpayX pages",
		Long:      ansi.ColoredBoldStatus(msg),
		Example: `RazorpayX wiki --list
				  RazorpayX wiki contact
  				  RazorpayX wiki fundAccount
  				  RazorpayX wiki payouts`,
		RunE: wc.runWikiCmd,
	}

	wc.cmd.Flags().Bool("list", false, "List all supported short cuts")
	wc.cmd.Flags().BoolVar(&wc.livemode, "live", false, "Receive live events (default: test)")

	return wc
}

func (oc *wikiCmd) runWikiCmd(cmd *cobra.Command, args []string) error {
	list, err := cmd.Flags().GetBool("list")
	if err != nil {
		return err
	}

	livemode, err := cmd.Flags().GetBool("live")
	if err != nil {
		return err
	}

	if list || len(args) == 0 {
		fmt.Println("wiki quickly opens RazorpayX pages. To use, run 'razorpayX wiki <shortcut>'.")
		fmt.Println("wiki supports the following shortcuts:")
		fmt.Println()

		shortcuts := wikiNames()
		sort.Strings(shortcuts)

		longest := getLongestShortcut(shortcuts)

		fmt.Printf("%s%s\n", padName("shortcut", longest), "    url")
		fmt.Printf("%s%s\n", padName("--------", longest), "    ---------")

		for _, shortcut := range shortcuts {
			maybeTestMode := ""
			if !livemode {
				maybeTestMode = "/test"
			}

			url := nameURLmap[shortcut]
			if strings.Contains(url, "%s") {
				url = fmt.Sprintf(url, maybeTestMode)
			}

			paddedName := padName(shortcut, longest)
			fmt.Printf("%s => %s\n", paddedName, url)
		}

		return nil
	}

	if url, ok := nameURLmap[args[0]]; ok {
		livemode, err := cmd.Flags().GetBool("live")
		if err != nil {
			return err
		}

		maybeTestMode := ""
		if !livemode {
			maybeTestMode = "/test"
		}

		if strings.Contains(url, "%s") {
			err = wiki.Browser(fmt.Sprintf(url, maybeTestMode))
		} else {
			err = wiki.Browser(url)
		}

		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Unsupported open command, given: %s", args[0])
	}

	return nil
}
