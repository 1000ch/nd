package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"

	"github.com/google/subcommands"
)

var semver *regexp.Regexp

func init() {
	semver = regexp.MustCompile(`v\d+.\d+.\d+`)
}

type remoteCommand struct {
}

func (*remoteCommand) Name() string {
	return "list-remote"
}

func (*remoteCommand) Synopsis() string {
	return "Show installable Node.js versions"
}

func (*remoteCommand) Usage() string {
	return `list-remote: Show installable Node.js versions`
}

func (i *remoteCommand) SetFlags(f *flag.FlagSet) {}

func (i *remoteCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	response, err := http.Get("https://nodejs.org/dist/")
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return subcommands.ExitFailure
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return subcommands.ExitFailure
	}

	versions := semver.FindAllString(string(bytes), -1)
	sort.Strings(versions)

	for _, v := range versions {
		fmt.Println(v)
	}

	return subcommands.ExitSuccess
}
