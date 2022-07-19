package append

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/machine-node-linker/manage-veneer/pkg/append"
)

func NewCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "append",
		Short: "create new semver file",
		Long:  "CLI to create and manage semvar veneer files for operator-framework/operator-registry",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			file, _ := cmd.Flags().GetString("file")

			if _, err := os.Stat(file); err != nil {
				return err
			}
			return nil
		},
		Args: cobra.NoArgs,
		RunE: append.Run,
	}

	cmd.Flags().String("bundle", "", "bundle image to add")
	cmd.Flags().String("channel", "candidate", "channel to add bundle to")
	cmd.Flags().Bool("no-lower", false, "dont add to channels below --channel")

	cmd.MarkFlagRequired("bundle")

	return cmd
}
