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
		Long:  "CLI to create and manage semvar veneer files for operator-framework/operator-registry",
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(
		appendcmd.NewCMD(),
		createcmd.NewCMD(),
	)

	rootCmd.PersistentFlags().String("file", "", "Semver Veneer File Path")
	rootCmd.MarkPersistentFlagRequired("file")

	return rootCmd
}
