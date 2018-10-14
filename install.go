package main

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/subcommands"
)

type installCommand struct{}

func (*installCommand) Name() string {
	return "install"
}

func (*installCommand) Synopsis() string {
	return "Install Node.js"
}

func (*installCommand) Usage() string {
	return `install <version>: Install Node.js <version>`
}

func (i *installCommand) SetFlags(f *flag.FlagSet) {}

func (i *installCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	args := f.Args()
	if len(args) != 1 {
		return subcommands.ExitFailure
	}

	version := normalizeVersion(args[0])
	platform := runtime.GOOS
	arch := normalizeArch(runtime.GOARCH)
	targetDir := filepath.Join(versionsDir, version)
	if err := prepareDir(targetDir); err != nil {
		return subcommands.ExitFailure
	}

	fileName := fmt.Sprintf("node-%s-%s-%s.tar.gz", version, platform, arch)
	url := fmt.Sprintf("https://nodejs.org/dist/%s/%s", version, fileName)

	if err := download(url, targetDir, fileName); err != nil {
		return subcommands.ExitFailure
	}

	if err := unarchive(targetDir, fileName); err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func download(url string, targetDir string, fileName string) error {
	distPath := filepath.Join(targetDir, fileName)

	file, err := os.Create(distPath)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		return err
	}

	response, err := http.Get(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func unarchive(targetDir string, fileName string) error {
	distPath := filepath.Join(targetDir, fileName)
	file, err := os.Open(distPath)
	if err != nil {
		return err
	}

	gzipReader, err := gzip.NewReader(file)
	if gzipReader != nil {
		defer gzipReader.Close()
	}
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(gzipReader)

	var header *tar.Header
	var targetPath string
	var fileMode os.FileMode
	var separator = string(os.PathSeparator)
	for {
		header, err = tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		pathList := strings.Split(header.Name, separator)[1:]
		targetPath = filepath.Join(targetDir, strings.Join(pathList, separator))
		fileMode = os.FileMode(header.Mode)

		switch header.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(targetPath, fileMode)
			if err != nil {
				return err
			}
		case tar.TypeReg:
			writer, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			io.Copy(writer, tarReader)

			err = os.Chmod(targetPath, fileMode)
			if err != nil {
				return err
			}

			writer.Close()
		default:
			fmt.Printf("Unable to untar type: %c in file %s", header.Typeflag, header.Name)
		}
	}

	if err := os.Remove(distPath); err != nil {
		return err
	}

	return nil
}
