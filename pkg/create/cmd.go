package create

import (
	"github.com/machine-node-linker/manage-veneer/pkg/semver"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, _ []string) error {
	file, _ := cmd.Flags().GetString("file")

	return semver.NewSemverVeneer().WriteFile(file)
}
