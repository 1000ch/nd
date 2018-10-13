package main

import (
	"os"
	"os/user"
)

func initialize() (err error) {
	user, err := user.Current()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(user.HomeDir+"/.nd", 0755); err != nil {
		return err
	}

	return nil
}
