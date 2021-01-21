package commands

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "RazorpayX",
	Short:   "A CLI to help you integrate RazorpayX with your application",
	Long: fmt.Sprintf("The command-line tool to interact with RazorpayX."),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init()  {
	rootCmd.AddCommand(newWikiCmd().cmd)
}