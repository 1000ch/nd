package main

import (
	"os"
	"os/user"
	"path/filepath"
)

func getBaseDir() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(user.HomeDir, ".nd"), nil
}

func getVersionsDir() (string, error) {
	baseDir, err := getBaseDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, "versions"), nil
}

func getBinaryDir() (string, error) {
	baseDir, err := getBaseDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, "bin"), nil
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
