package commands

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"

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

	stripeAccount string
}

func newFixturesCmd(cfg *config.Config) *FixturesCmd {
	fixturesCmd := &FixturesCmd{
		Cfg: cfg,
	}

	fixturesCmd.Cmd = &cobra.Command{
		Use:   "fixtures",
		Args:  validators.ExactArgs(1),
		Short: "Run fixtures to populate your account with data",
		Long:  `Run fixtures to populate your account with data`,
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

	fixture, err := fixtures.NewFixture(
		afero.NewOsFs(),
		apiKey,
		apiSecret,
		client.DefaultAPIBaseURL,
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
