package commands

import (
	"fmt"
	"github.com/rzp-gt/razorpayx-cli/internal/ansi"
	"github.com/rzp-gt/razorpayx-cli/internal/client"
	"github.com/rzp-gt/razorpayx-cli/internal/fixtures"
	"github.com/rzp-gt/razorpayx-cli/internal/validators"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type triggerCmd struct {
	cmd *cobra.Command

	fs         afero.Fs
	apiBaseURL string
}

func newTriggerCmd() *triggerCmd {
	tc := &triggerCmd{}
	tc.fs = afero.NewOsFs()

	msg := fmt.Sprintf("Trigger specific webhook events to be sent. Webhooks events created through \n"+
		"the trigger command will also create all necessary side-effect events that are \n"+
		"needed to create the triggered event as well as the corresponding API objects. \n \n"+
		"%s \n"+
		"%s \n", "Supported events:",
		fixtures.EventList())

	tc.cmd = &cobra.Command{
		Use:       "trigger <event>",
		Args:      validators.MaximumNArgs(1),
		ValidArgs: fixtures.EventNames(),
		Short:     "Trigger test webhook events",
		Long:      ansi.ColoredBoldStatus(msg),
		Example:   `RazorpayX trigger payout.created`,
		RunE:      tc.runTriggerCmd,
	}

	// Hidden configuration flags, useful for dev/debugging
	tc.cmd.Flags().StringVar(&tc.apiBaseURL, "api-base", client.DefaultAPIBaseURL, "Sets the API base URL")
	tc.cmd.Flags().MarkHidden("api-base") // #nosec G104

	return tc
}

func (tc *triggerCmd) runTriggerCmd(cmd *cobra.Command, args []string) error {

	apiKey, err := Config.Profile.GetAPIKey(false)
	if err != nil {
		return err
	}

	apiSecret, err := Config.Profile.GetAPISecret(false)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		cmd.Help()

		return nil
	}

	event := args[0]

	var fixture *fixtures.Fixture
	if file, ok := fixtures.Events[event]; ok {
		fixture, err = fixtures.BuildFromFixture(tc.fs, apiKey, apiSecret, tc.apiBaseURL, file)
		if err != nil {
			return err
		}
	} else {
		exists, _ := afero.Exists(tc.fs, event)
		if !exists {
			return fmt.Errorf(fmt.Sprintf("event %s is not supported.", event))
		}

		fixture, err = fixtures.BuildFromFixture(tc.fs, apiKey, apiSecret, tc.apiBaseURL, event)
		if err != nil {
			return err
		}
	}

	err = fixture.Execute()
	if err == nil {
		fmt.Println("Trigger succeeded!")
	} else {
		fmt.Printf("Trigger failed: %s\n", err)
	}

	return err
}
