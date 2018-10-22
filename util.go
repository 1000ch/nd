package main

import (
	"os"
	"sort"
	"strings"

	repo "github.com/1000ch/nd/repository"
)

func unique(args []string) []string {
	versions := []string{}

	m := make(map[string]bool)
	for _, v := range args {
		if !m[v] {
			m[v] = true
			versions = append(versions, v)
		}
	}

	return versions
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

func normalizeVersions(args []string) []*repo.Version {
	versions := unique(args)
	semvers := make([]*repo.Version, len(versions))
	for i, version := range versions {
		semvers[i] = repo.NewVersion(version)
	}

	sort.Sort(repo.Versions(semvers))

	return semvers
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
