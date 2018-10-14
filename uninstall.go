package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
)

type uninstallCommand struct{}

func (*uninstallCommand) Name() string {
	return "uninstall"
}

func (*uninstallCommand) Synopsis() string {
	return "Uninstall Node.js"
}

func (*uninstallCommand) Usage() string {
	return `uninstall <version>: Uninstall Node.js.`
}

func (i *uninstallCommand) SetFlags(f *flag.FlagSet) {}

func (i *uninstallCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	err := initialize()
	if err != nil {
		return subcommands.ExitFailure
	}

	args := f.Args()
	if len(args) != 1 {
		return subcommands.ExitFailure
	}

	version := normalizeVersion(args[0])
	versionsDir, err := getVersionsDir()
	if err != nil {
		return subcommands.ExitFailure
	}

	targetDir := filepath.Join(versionsDir, version)
	if err := prepareDir(targetDir); err != nil {
		return subcommands.ExitFailure
	}

	if err := os.RemoveAll(targetDir); err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
