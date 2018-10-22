package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/google/subcommands"
)

var sv *regexp.Regexp

func init() {
	sv = regexp.MustCompile(`v\d+.\d+.\d+`)
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

	versions := sv.FindAllString(string(bytes), -1)
	semvers := normalizeVersions(versions)

	var count int32
	m1 := make(map[string]bool)
	m2 := make(map[int64]bool)
	for _, v := range semvers {
		major := v.Version.Major()
		minor := v.Version.Minor()
		version := fmt.Sprintf("%d.%d", major, minor)
		if major == 0 && m1[version] == false {
			m1[version] = true
			count = 0
			fmt.Println()
			fmt.Println()
		} else if major != 0 && m2[major] == false {
			m2[major] = true
			count = 0
			fmt.Println()
			fmt.Println()
		}

		count++
		if count%8 == 0 {
			fmt.Printf("%v\n", v)
		} else {
			fmt.Printf("%-10v", v)
		}
	}

	return subcommands.ExitSuccess
}
