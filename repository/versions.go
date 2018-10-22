package repository

// Versions is a collection of Version instances and implements the sort
// interface. See the sort package for more details.
// https://golang.org/pkg/sort/
type Versions []*Version

// Len returns the length of a collection. The number of Version instances
// on the slice.
func (vs Versions) Len() int {
	return len(vs)
}

// Less is needed for the sort interface to compare two Version objects on the
// slice. If checks if one is less than the other.
func (vs Versions) Less(i, j int) bool {
	return vs[i].LessThan(vs[j])
}

// Swap is needed for the sort interface to replace the Version objects
// at two different positions in the slice.
func (vs Versions) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}
