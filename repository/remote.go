package repository

import (
	"fmt"
)

type Remote struct {
	Platform string
	Arch     string
}

func (r *Remote) Filename(v *Version) string {
	return fmt.Sprintf("node-%s-%s-%s.tar.gz", v.String(), r.Platform, r.Arch)
}

func (r *Remote) URL(v *Version) string {
	return fmt.Sprintf("https://nodejs.org/dist/%s/%s", v.String(), r.Filename(v))
}
