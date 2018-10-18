package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
)

type useCommand struct {
	source bool
}

func (*useCommand) Name() string {
	return "use"
}

func (*useCommand) Synopsis() string {
	return "Activate specified Node.js"
}

func (*useCommand) Usage() string {
	return `use <version>: Activate Node.js <version>`
}

func (i *useCommand) SetFlags(f *flag.FlagSet) {}

func (i *useCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	args := f.Args()
	if len(args) != 1 {
		return subcommands.ExitFailure
	}

	version := normalizeVersion(args[0]).String()
	symlinkTarget := filepath.Join(versionsDir, version, "bin", "node")
	symlinkPath := filepath.Join(binaryDir, "node")

	if err := os.Remove(symlinkPath); err != nil {
		return subcommands.ExitFailure
	}

	if err := os.Symlink(symlinkTarget, symlinkPath); err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
