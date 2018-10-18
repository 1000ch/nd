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
	files, err := ioutil.ReadDir(versionsDir)
	if err != nil {
		return subcommands.ExitFailure
	}

	versions := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			versions = append(versions, file.Name())
		}
	}

	semvers := normalizeVersions(versions)
	for _, v := range semvers {
		fmt.Println(v)
	}

	return subcommands.ExitSuccess
}
