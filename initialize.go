package main

func initialize() error {
	baseDir, err := getBaseDir()
	if err != nil {
		return err
	}

	if err := prepareDir(baseDir); err != nil {
		return err
	}

	return nil
}
