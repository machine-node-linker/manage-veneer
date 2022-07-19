package append

import (
	"fmt"

	"github.com/machine-node-linker/manage-veneer/pkg/semver"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, _ []string) error {
	file, _ := cmd.Flags().GetString("file")
	bundle, _ := cmd.Flags().GetString("bundle")
	channel, _ := cmd.Flags().GetString("channel")
	addLower, _ := cmd.Flags().GetBool("add-lower")

	sv, err := semver.LoadFile(file)
	if err != nil {
		return fmt.Errorf("Unable to load file: %w", err)
	}

	var channels []string

	if addLower {
		channels = semver.GetIncludedChannels(channel)
	} else {
		channels[0] = channel
	}

	for _, ch := range channels {
		if err := sv.AddBundleToChannel(bundle, ch); err != nil {
			return fmt.Errorf("Unable to append to channel: %w", err)
		}
	}

	if err = sv.WriteFile(file); err != nil {
		return fmt.Errorf("Unable to write semver file: %w", err)
	}

	return nil
}
