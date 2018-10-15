package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/golang/glog"
	"github.com/mitchellh/go-homedir"
)

func getBaseDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		glog.Error(err)
	}

	return filepath.Join(dir, ".nd")
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
	number1, err1 := regexp.Compile(`\d+.\d+.\d+`)
	number2, err2 := regexp.Compile(`\d+.\d+`)

	if err1 != nil || err2 != nil {
		return version
	}

	if number1.MatchString(version) {
		return fmt.Sprintf("v%s", version)
	}

	if number2.MatchString(version) {
		return fmt.Sprintf("v%s.0", version)
	}

	return version
}

func normalizeArch(goarch string) string {
	if strings.Contains(goarch, "arm64") ||
		strings.Contains(goarch, "amd64") {
		return "x64"
	}

	if strings.Contains(goarch, "386") ||
		strings.Contains(goarch, "arm") ||
		strings.Contains(goarch, "amd64p32") {
		return "x86"
	}

	return ""
}
