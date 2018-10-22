package repository

import (
	"fmt"
)

type Remote struct {
	Platform string
	Arch     string
}

func (r *Remote) Filename(v *Version) string {
	format := "node-%s-%s-%s.tar.gz"

	return fmt.Sprintf(format, v.String(), r.Platform, r.Arch)
}

func (r *Remote) Url(v *Version) string {
	format := "https://nodejs.org/dist/%s/%s"

	return fmt.Sprintf(format, v.String(), r.Filename(v))
}
