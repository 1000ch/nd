package main

import (
	"context"
	"flag"
	"os"

	"github.com/1000ch/nd/repo"
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
	return `uninstall <version>: Uninstall Node.js <version>`
}

func (i *uninstallCommand) SetFlags(f *flag.FlagSet) {}

func (i *uninstallCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	v := repo.NewVersion(f.Args()[0])

	if err := prepareDir(local.NodeDir(v)); err != nil {
		return subcommands.ExitFailure
	}

	if err := os.RemoveAll(local.NodeDir(v)); err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
