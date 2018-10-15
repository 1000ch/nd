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
	"time"

	"github.com/briandowns/spinner"
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

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Downloading Node.js %s", version)
	s.Start()
	if err := download(url, targetDir, fileName); err != nil {
		return subcommands.ExitFailure
	}
	s.Stop()

	return subcommands.ExitSuccess
}

func download(url string, targetDir string, fileName string) error {
	response, err := http.Get(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	gzipReader, err := gzip.NewReader(response.Body)
	if gzipReader != nil {
		defer gzipReader.Close()
	}
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		separator := string(os.PathSeparator)
		pathList := strings.Split(header.Name, separator)[1:]
		targetPath := filepath.Join(targetDir, strings.Join(pathList, separator))
		fileMode := os.FileMode(header.Mode)

		switch header.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(targetPath, fileMode)
			if err != nil {
				return err
			}
		case tar.TypeReg, tar.TypeLink, tar.TypeSymlink:
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

	return nil
}
