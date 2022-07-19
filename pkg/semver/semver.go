package semver

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type semverVeneerBundleEntry struct {
	Image string
}

type bundleSlice struct {
	Bundles []semverVeneerBundleEntry `yaml:"Bundles"`
}

func (b *bundleSlice) add(i string) error {
	if !b.contains(i) {
		b.Bundles = append(b.Bundles, semverVeneerBundleEntry{i})
	}

	return nil
}

func (b *bundleSlice) contains(i string) bool {
	for _, entry := range b.Bundles {
		if entry.Image == i {
			return true
		}
	}

	return false
}

type semverVeneer struct {
	Schema                string      `yaml:"Schema"`
	GenerateMajorChannels bool        `yaml:"GenerateMajorChannels"`
	GenerateMinorChannels bool        `yaml:"GenerateMinorChannels"`
	AvoidSkipPatch        bool        `yaml:"AvoidSkipPatch,omitempty"`
	Candidate             bundleSlice `yaml:"Candidate"`
	Fast                  bundleSlice `yaml:"Fast"`
	Stable                bundleSlice `yaml:"Stable"`
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

func GetIncludedChannels(ch string) []string {
	channels := make([]string, 3)
	for i, chs := range getChannelOrder() {
		channels[i] = string(chs)
	}

	for range channels {
		if channels[0] == ch {
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

	return svch.add(bundle)
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
