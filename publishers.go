package oops

const (
	InMemory = "in-memory"
	Disk     = "disk"
)

// Publisher is the componet in charge of persisting the oops.
type Publisher interface {
	Write(o Oops) error
	Read(id string) (*Oops, error)
}
