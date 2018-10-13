package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type installCommand struct{}

func (*installCommand) Name() string {
	return "install"
}

func (*installCommand) Synopsis() string {
	return "Install Node.js"
}

func (*installCommand) Usage() string {
	return `install <version>: Install Node.js.`
}

func (i *installCommand) SetFlags(f *flag.FlagSet) {}

func (i *installCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	err := initialize()
	if err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
