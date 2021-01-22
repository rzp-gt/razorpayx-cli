package commands

import (
	"github.com/rzp-gt/razorpayx-cli/internal/ansi"
	"github.com/spf13/cobra"

	"github.com/rzp-gt/razorpayx-cli/internal/config"
)

type configCmd struct {
	cmd    *cobra.Command
	config *config.Config

	list   bool
	edit   bool
	unset  string
	set    bool
	create bool
}

func newConfigCmd() *configCmd {
	cc := &configCmd{
		config: &Config,
	}
	msg := "config lets you set and unset specific configuration values or your profile \n" +
		"if you need more granular control over the configuration.\n"

	cc.cmd = &cobra.Command{
		Use:   "config",
		Short: "Manually change the config values for the CLI",
		Long: ansi.ColoredBoldStatus(msg),
		Example: `razorpayx config --list`,
		RunE:    cc.runConfigCmd,
	}

	cc.cmd.Flags().BoolVar(&cc.list, "list", false, "List configs")
	cc.cmd.Flags().BoolVarP(&cc.edit, "edit", "e", false, "Open an editor to the config file")
	cc.cmd.Flags().StringVar(&cc.unset, "unset", "", "Unset a specific config field")
	cc.cmd.Flags().BoolVar(&cc.set, "set", false, "Set a config field to some value")
	cc.cmd.Flags().BoolVar(&cc.create, "create", false, "create a config profile")

	cc.cmd.Flags().SetInterspersed(false) // allow args to happen after flags to enable 2 arguments to --set

	return cc
}

func (cc *configCmd) runConfigCmd(cmd *cobra.Command, args []string) error {
	switch ok := true; ok {
	case cc.set && len(args) == 2:
		return cc.config.Profile.WriteConfigField(args[0], args[1])
	case cc.unset != "":
		return cc.config.Profile.DeleteConfigField(cc.unset)
	case cc.list:
		return cc.config.PrintConfig()
	case cc.edit:
		return cc.config.EditConfig()
	case cc.create:
		return cc.config.CreateProfile()
	default:
		// no flags set or unrecognized flags/args
		return cc.cmd.Help()
	}
}
