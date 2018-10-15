package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/google/subcommands"
	"github.com/mitchellh/go-homedir"
)

var version string
var baseDir string
var binaryDir string
var versionsDir string

func init() {
	homeDir, err := homedir.Dir()
	if err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	baseDir = filepath.Join(homeDir, ".nd")
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
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(subcommands.Alias("ls", &listCommand{}), "")
	subcommands.Register(&installCommand{}, "")
	subcommands.Register(&uninstallCommand{}, "")
	subcommands.Register(&useCommand{}, "")
	subcommands.Register(&listCommand{}, "")
	subcommands.Register(&versionCommand{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
