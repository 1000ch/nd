package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/google/subcommands"
)

type listCommand struct {
}

func (*listCommand) Name() string {
	return "list"
}

func (*listCommand) Synopsis() string {
	return "Show installed Node.js versions"
}

func (*listCommand) Usage() string {
	return `list: Show installed Node.js versions`
}

func (i *listCommand) SetFlags(f *flag.FlagSet) {}

func (i *listCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	files, err := ioutil.ReadDir(local.VersionsDir())
	if err != nil {
		return subcommands.ExitFailure
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
		}
	}

	return subcommands.ExitSuccess
}
