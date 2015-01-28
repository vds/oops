package publishers

import "github.com/vds/oops"

const (
	InMemory = "in-memory"
	Disk     = "disk"
)

// Publisher is the componet in charge of persisting the oops.
type Publisher interface {
	Write(o oops.Oops) error
	Read(id string) (*oops.Oops, error)
}
