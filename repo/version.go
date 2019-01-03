package repo

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/golang/glog"
)

type Version struct {
	Version *semver.Version
}

func NewVersion(arg string) *Version {
	version, err := semver.NewVersion(arg)
	if err != nil {
		glog.Errorf("Error parsing version: %s", err)
	}

	return &Version{version}
}

func (v *Version) String() string {
	version := v.Version.String()

	if version[:1] != "v" {
		return fmt.Sprintf("v%s", version)
	}

	return version
}

func (v *Version) LessThan(o *Version) bool {
	return v.Version.Compare(o.Version) < 0
}
