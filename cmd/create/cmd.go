package root

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/machine-node-linker/manage-veneer/pkg/create"
)

func NewCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create new semver file",
		Long:  "CLI to create and manage semvar veneer files for operator-framework/operator-registry",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				return fmt.Errorf("unable to read file name: %w", err)
			}
			if overwrite, _ := cmd.Flags().GetBool("overwrite"); !overwrite {
				_, err = os.Stat(file)
				if !os.IsNotExist(err) {
					return fmt.Errorf("will not modify %s without --overwrite", file)
				}
			}
			dir, _ := path.Split(file)
			if mkdir, _ := cmd.Flags().GetBool("make-dirs"); mkdir {
				if dir != "" {
					if err := os.MkdirAll(dir, 0755); err != nil {
						return fmt.Errorf("unable to make path: %w", err)
					}
				}
			} else {
				stat, err := os.Stat(dir)
				if !os.IsNotExist(err) {
					return fmt.Errorf("directory %s does not exist", dir)
				}
				if !stat.IsDir() {
					return fmt.Errorf("%s is not a Directory", dir)
				}
			}
			return nil
		},
		Args: cobra.NoArgs,
		RunE: create.Run,
	}

	cmd.Flags().Bool("overwrite", false, "overwrite existing file")
	cmd.Flags().Bool("make-dirs", true, "make missing directories in file path")

	return cmd
}
