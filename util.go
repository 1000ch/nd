package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var semver1, semver2 *regexp.Regexp

func init() {
	semver1 = regexp.MustCompile(`\d+.\d+.\d+`)
	semver2 = regexp.MustCompile(`\d+.\d+`)
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
	if semver1.MatchString(version) {
		return fmt.Sprintf("v%s", version)
	}

	if semver2.MatchString(version) {
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
