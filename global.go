package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
)

type globalCommand struct {
	source bool
}

func (*globalCommand) Name() string {
	return "global"
}

func (*globalCommand) Synopsis() string {
	return "Set specified version to global"
}

func (*globalCommand) Usage() string {
	return `global <version>: Set specified version to global.`
}

func (i *globalCommand) SetFlags(f *flag.FlagSet) {}

func (i *globalCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	args := f.Args()
	if len(args) != 1 {
		return subcommands.ExitFailure
	}

	version := normalizeVersion(args[0])
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
