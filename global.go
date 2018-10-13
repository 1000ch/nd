package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type globalCommand struct {
	source bool
}

func (*globalCommand) Name() string {
	return "gloabl"
}

func (*globalCommand) Synopsis() string {
	return "Set specified version to global"
}

func (*globalCommand) Usage() string {
	return `global <version>: Set specified version to global.`
}

func (i *globalCommand) SetFlags(f *flag.FlagSet) {}

func (i *globalCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	err := initialize()
	if err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
