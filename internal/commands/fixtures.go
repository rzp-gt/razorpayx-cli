package commands

import (
	"github.com/rzp-gt/razorpayx-cli/internal/ansi"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"strings"

	"github.com/rzp-gt/razorpayx-cli/internal/client"
	"github.com/rzp-gt/razorpayx-cli/internal/config"
	"github.com/rzp-gt/razorpayx-cli/internal/fixtures"
	"github.com/rzp-gt/razorpayx-cli/internal/validators"
)

// FixturesCmd prints a list of all the available sample projects that users can
// generate
type FixturesCmd struct {
	Cmd *cobra.Command
	Cfg *config.Config
}

func newFixturesCmd(cfg *config.Config) *FixturesCmd {
	fixturesCmd := &FixturesCmd{
		Cfg: cfg,
	}

	fixturesCmd.Cmd = &cobra.Command{
		Use:   "fixtures",
		Args:  validators.ExactArgs(1),
		Short: "Run fixtures to populate your account with data",
		Long:  ansi.ColoredBoldStatus("Run fixtures to define workflows for APIs"),
		RunE:  fixturesCmd.runFixturesCmd,
	}

	return fixturesCmd
}

func (fc *FixturesCmd) runFixturesCmd(cmd *cobra.Command, args []string) error {

	apiKey, err := fc.Cfg.Profile.GetAPIKey(false)
	if err != nil {
		return err
	}

	apiSecret, err := fc.Cfg.Profile.GetAPISecret(false)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return nil
	}

	//checklist --api-key rzp_test_1DP5mmOlF5G5ag --api-secret thisissupersecret
	res := strings.Contains(args[0], "checklist")

	url := client.DefaultAPIBaseURL
	if res {
		url = "http://192.168.0.103:8000"
		apiKey = "rzp_test_1DP5mmOlF5G5ag"
		apiSecret = "thisissupersecret"
	}

	fixture, err := fixtures.NewFixture(
		afero.NewOsFs(),
		apiKey,
		apiSecret,
		url,
		args[0],
	)
	if err != nil {
		return err
	}

	err = fixture.Execute()
	if err != nil {
		return err
	}

	return nil
}
