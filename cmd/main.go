package main

import (
	"os"

	"github.com/machine-node-linker/manage-veneer/cmd/root"
	"github.com/machine-node-linker/manage-veneer/pkg/github"
)

func main() {
	cmd := root.NewCMD()
	cmd.SetErr(github.ErrorWriter{})

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
