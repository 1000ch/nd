package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"

	"github.com/golang/glog"
)

func getBaseDir() string {
	user, err := user.Current()
	if err != nil {
		glog.Error(err)
	}

	return filepath.Join(user.HomeDir, ".nd")
}

func prepareDir(p string) error {
	if _, err := os.Stat(p); !os.IsNotExist(err) {
		return nil
	}

	if err := os.MkdirAll(p, 0755); err != nil {
		return err
	}

	return nil
}

func normalizeVersion(version string) string {
	normalized := regexp.MustCompile(`v\d+.\d+.\d+`)

	if normalized.MatchString(version) {
		return version
	}

	number1 := regexp.MustCompile(`\d+.\d+.\d+`)
	number2 := regexp.MustCompile(`\d+.\d+`)

	if number1.MatchString(version) {
		return fmt.Sprintf("v%s", version)
	}

	if number2.MatchString(version) {
		return fmt.Sprintf("v%s.0", version)
	}

	return version
}

func normalizeArch(goarch string) string {
	x64 := regexp.MustCompile(`amd64|arm64`)
	x86 := regexp.MustCompile(`386|arm|amd64p32`)

	if x64.MatchString(goarch) {
		return "x64"
	}

	if x86.MatchString(goarch) {
		return "x86"
	}

	return ""
}
