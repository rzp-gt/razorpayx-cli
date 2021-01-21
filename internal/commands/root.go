package commands

import (
	"fmt"
	"github.com/rzp-gt/razorpayx-cli/internal/config"
	"github.com/spf13/cobra"
	"os"
)

// Config is the cli configuration for the user
var Config config.Config

var rootCmd = &cobra.Command{
	Use:   "RazorpayX",
	Short: "A CLI to help you integrate RazorpayX with your application",
	Long:  fmt.Sprintf("The command-line tool to interact with RazorpayX."),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(Config.InitConfig)
	rootCmd.PersistentFlags().StringVar(&Config.Profile.APIKey, "api-key", "", "Your API key to use for the command")
	rootCmd.PersistentFlags().StringVar(&Config.Profile.APISecret, "api-secret", "", "Your API secret to use for the command")
	rootCmd.PersistentFlags().StringVar(&Config.ProfilesFile, "config", "", "config file (default is $HOME/.config/razorpayx/config.toml)")
	rootCmd.PersistentFlags().StringVar(&Config.Profile.DeviceName, "device-name", "", "device name")
	rootCmd.PersistentFlags().StringVarP(&Config.Profile.ProfileName, "project-name", "p", "default", "the project name to read from for config")

	rootCmd.AddCommand(newConfigCmd().cmd)
	rootCmd.AddCommand(newWikiCmd().cmd)
	rootCmd.AddCommand(newGetCmd().reqs.Cmd)
	rootCmd.AddCommand(newPostCmd().reqs.Cmd)
}
