package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type versionCommand struct{}

func (*versionCommand) Name() string {
	return "version"
}

func (*versionCommand) Synopsis() string {
	return "Show nd version"
}

func (*versionCommand) Usage() string {
	return `version: Show nd version`
}

func (i *versionCommand) SetFlags(f *flag.FlagSet) {}

func (i *versionCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println(version)

	return subcommands.ExitSuccess
}
