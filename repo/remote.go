package repo

import (
	"fmt"
)

// Remote is a struct for remote archives
type Remote struct {
	Platform string
	Arch     string
}

// Filename returns an archive file name from version
func (r *Remote) Filename(v *Version) string {
	return fmt.Sprintf("node-%s-%s-%s.tar.gz", v.String(), r.Platform, r.Arch)
}

// URL returns an archive URL
func (r *Remote) URL(v *Version) string {
	return fmt.Sprintf("https://nodejs.org/dist/%s/%s", v.String(), r.Filename(v))
}
