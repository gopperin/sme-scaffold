package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// StartCmd version
var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "get version info",
		Example: "sme-scaffold version",
		PreRun: func(cmd *cobra.Command, args []string) {
			// pre不需要做操作
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println("v2021.07.04")
	return nil
}
