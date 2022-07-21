package root

import (
	"github.com/spf13/cobra"

	appendcmd "github.com/machine-node-linker/manage-veneer/cmd/append"
	createcmd "github.com/machine-node-linker/manage-veneer/cmd/create"
)

func NewCMD() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "manage-veneer",
		Short: "semver veneer tool",
		Long:  "CLI to create and manage semver veneer files for operator-framework/operator-registry",
		Args:  cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}

	rootCmd.AddCommand(
		appendcmd.NewCMD(),
		createcmd.NewCMD(),
	)

	return rootCmd
}
