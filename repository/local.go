package repository

import (
	"path/filepath"
)

type Local struct {
	HomeDir string
}

func (l *Local) BaseDir() string {
	return filepath.Join(l.HomeDir, ".nd")
}

func (l *Local) VersionsDir() string {
	return filepath.Join(l.BaseDir(), "versions")
}

func (l *Local) NodeDir(v *Version) string {
	return filepath.Join(l.VersionsDir(), v.String())
}

func (l *Local) NodePath(v *Version) string {
	return filepath.Join(l.NodeDir(v), "bin", "node")
}

func (l *Local) BinDir() string {
	return filepath.Join(l.BaseDir(), "bin")
}

func (l *Local) BinPath() string {
	return filepath.Join(l.BinDir(), "node")
}
