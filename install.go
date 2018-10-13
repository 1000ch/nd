package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type installCommand struct {
	source bool
}

func (*installCommand) Name() string {
	return "install"
}

func (*installCommand) Synopsis() string {
	return "Install Node.js"
}

func (*installCommand) Usage() string {
	return `install [-source] <version>: Install Node.js.`
}

func (i *installCommand) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&i.source, "source", false, "compile from source")
}

func (i *installCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
