package append

import (
	"github.com/machine-node-linker/manage-veneer/pkg/semver"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, _ []string) error {
	file, _ := cmd.Flags().GetString("file")
	bundle, _ := cmd.Flags().GetString("bundle")
	channel, _ := cmd.Flags().GetStringSlice("channel")
	no_lower, _ := cmd.Flags().GetBool("no-lower")

	sv, err := semver.LoadFile(file)
	if err != nil {
		return err
	}
	var channels []string
	if no_lower {
		channels = channel
	} else {
		channels = semver.GetIncludedChannels(channel)
	}

	for _, ch := range channels {
		if err := sv.AddBundleToChannel(bundle, ch); err != nil {
			return err
		}
	}

	return sv.WriteFile(file)
}
