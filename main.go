package main

import (
	"context"
	"flag"
	"os"
	"runtime"

	"github.com/1000ch/nd/repo"
	"github.com/golang/glog"
	"github.com/google/subcommands"
	"github.com/mitchellh/go-homedir"
)

var version string
var local repo.Local
var remote repo.Remote

func init() {
	homeDir, err := homedir.Dir()
	if err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	local = repo.Local{homeDir}
	remote = repo.Remote{runtime.GOOS, normalizeArch(runtime.GOARCH)}

	if err := prepareDir(local.BaseDir()); err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	if err := prepareDir(local.BinDir()); err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	if err := prepareDir(local.VersionsDir()); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(subcommands.Alias("ls", &listCommand{}), "")
	subcommands.Register(subcommands.Alias("ls-remote", &remoteCommand{}), "")
	subcommands.Register(&installCommand{}, "")
	subcommands.Register(&uninstallCommand{}, "")
	subcommands.Register(&useCommand{}, "")
	subcommands.Register(&listCommand{}, "")
	subcommands.Register(&remoteCommand{}, "")
	subcommands.Register(&versionCommand{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
