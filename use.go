package main

import (
	"context"
	"flag"
	"os"

	"github.com/1000ch/nd/repo"
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
	v := repo.NewVersion(f.Args()[0])

	if err := os.Remove(local.BinPath()); err != nil {
		return subcommands.ExitFailure
	}

	if err := os.Symlink(local.NodePath(v), local.BinPath()); err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
