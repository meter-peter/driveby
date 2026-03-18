package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Set via ldflags at build time:
//
//	go build -ldflags "-X github.com/meter-peter/driveby/driveby-cli/internal/cli.version=1.0.0"
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the DriveBy version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("driveby %s (commit: %s, built: %s)\n", version, commit, date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
