package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"sme-stage/cmd/api"
	"sme-stage/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "sme-stage",
	Short:        "sme-stage",
	SilenceUsage: true,
	Long:         `sme-stage`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	fmt.Println("欢迎使用sme-stage")
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
}

// Execute Execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
