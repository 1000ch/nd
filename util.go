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
	number3 := regexp.MustCompile(`\d+`)

	if number1.MatchString(version) {
		return fmt.Sprintf("v%s", version)
	}

	if number2.MatchString(version) {
		return fmt.Sprintf("v%s.0", version)
	}

	if number3.MatchString(version) {
		return fmt.Sprintf("v%s.0.0", version)
	}

	return version
}

func normalizeArch(goarch string) string {
	return "x64"
}
