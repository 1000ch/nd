package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

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

	dirs := filter(files, func(fi os.FileInfo) bool {
		return fi.IsDir()
	})

	for _, d := range dirs {
		fmt.Println(d.Name())
	}

	return subcommands.ExitSuccess
}

func filter(fis []os.FileInfo, f func(os.FileInfo) bool) []os.FileInfo {
	filtered := make([]os.FileInfo, 0)

	for _, fi := range fis {
		if f(fi) {
			filtered = append(filtered, fi)
		}
	}

	return filtered
}
