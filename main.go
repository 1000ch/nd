package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/google/subcommands"
)

var baseDir string
var binaryDir string
var versionsDir string

func main() {
	baseDir = getBaseDir()
	binaryDir = filepath.Join(baseDir, "bin")
	versionsDir = filepath.Join(baseDir, "versions")

	if err := prepareDir(baseDir); err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	if err := prepareDir(binaryDir); err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	if err := prepareDir(versionsDir); err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&installCommand{}, "")
	subcommands.Register(&uninstallCommand{}, "")
	subcommands.Register(&useCommand{}, "")
	subcommands.Register(&versionsCommand{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
