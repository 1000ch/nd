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
	"strings"
	"time"

	"github.com/1000ch/nd/repo"
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
	v := repo.NewVersion(f.Args()[0])
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Downloading Node.js %s", v.String())

	if err := prepareDir(local.NodeDir(v)); err != nil {
		return subcommands.ExitFailure
	}

	s.Start()
	if err := download(remote.URL(v), local.NodeDir(v)); err != nil {
		return subcommands.ExitFailure
	}
	s.Stop()

	return subcommands.ExitSuccess
}

func download(url string, targetDir string) error {
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
