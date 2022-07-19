package semver

import (
	"fmt"
	"os"
	"sort"

	"sigs.k8s.io/yaml"
)

type semverVeneerBundleEntry struct {
	Image string
}

type bundleSlice struct {
	Bundles []semverVeneerBundleEntry
}

type semverVeneer struct {
	Schema                string
	GenerateMajorChannels bool
	GenerateMinorChannels bool
	AvoidSkipPatch        bool `json:",omitempty"`
	Candidate             bundleSlice
	Fast                  bundleSlice
	Stable                bundleSlice
}

type channel string

const (
	candidateChannel channel = "candidate"
	fastChannel      channel = "fast"
	stableChannel    channel = "stable"
)

func getChannelOrder() []channel {
	return []channel{stableChannel, fastChannel, candidateChannel}
}

func GetIncludedChannels(ch []string) []string {
	channels := make([]string, 3)
	for i, chs := range getChannelOrder() {
		channels[i] = string(chs)
	}

	if len(ch) > 0 {
		sort.Slice(ch, func(i, j int) bool {
			switch channel(ch[i]) {
			case candidateChannel:
				return true
			case stableChannel:
				return false
			}
			return channel(ch[j]) == stableChannel
		})
	}

	for range channels {
		if channels[0] == ch[0] {
			return channels
		}
		channels = channels[1:]
	}
	return channels
}

func LoadFile(filename string) (*semverVeneer, error) {
	sv := &semverVeneer{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, sv); err != nil {
		return nil, err
	}

	return sv, nil
}

func (sv *semverVeneer) AddBundleToChannel(bundle string, ch string) error {
	svch, err := sv.getChannel(channel(ch))
	if err != nil {
		return err
	}
	svch.Bundles = append(svch.Bundles, semverVeneerBundleEntry{bundle})
	return nil
}

func (sv *semverVeneer) getChannel(ch channel) (*bundleSlice, error) {
	// fix this
	switch ch {
	case candidateChannel:
		return &sv.Candidate, nil
	case fastChannel:
		return &sv.Fast, nil
	case stableChannel:
		return &sv.Stable, nil
	}
	return nil, fmt.Errorf("invalid channel %s", ch)
}

func (sv *semverVeneer) WriteFile(filename string) error {
	data, err := yaml.Marshal(sv)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func NewSemverVeneer() *semverVeneer {
	return &semverVeneer{
		GenerateMajorChannels: true,
		GenerateMinorChannels: false,
		Schema:                "olm.semver",
	}
}
