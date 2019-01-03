package repo

import (
	"path/filepath"
)

// Local is a struct to manage local directory structure
type Local struct {
	HomeDir string
}

// BaseDir returns base directory path
func (l *Local) BaseDir() string {
	return filepath.Join(l.HomeDir, ".nd")
}

// VersionsDir returns dist directory path
func (l *Local) VersionsDir() string {
	return filepath.Join(l.BaseDir(), "versions")
}

// NodeDir returns node directory path
func (l *Local) NodeDir(v *Version) string {
	return filepath.Join(l.VersionsDir(), v.String())
}

// NodePath returns node binary path
func (l *Local) NodePath(v *Version) string {
	return filepath.Join(l.NodeDir(v), "bin", "node")
}

// BinDir returns binary directory path
func (l *Local) BinDir() string {
	return filepath.Join(l.BaseDir(), "bin")
}

// BinPath returns binary path
func (l *Local) BinPath() string {
	return filepath.Join(l.BinDir(), "node")
}
