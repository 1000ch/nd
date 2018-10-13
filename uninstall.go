package main

import (
	"context"
	"flag"

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

	return subcommands.ExitSuccess
}
