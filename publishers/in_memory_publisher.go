package publishers

import (
	"errors"

	"github.com/vds/oops"
)

type InMemoryOopsStorage map[string]*oops.Oops

// InMemoryPublisher is a simple publisher that stores oopses in memory, convenient for testing.
type InMemoryPublisher struct {
	Storage InMemoryOopsStorage
}

// Write writes the binary marshalling of a oops on the disk.
func (p InMemoryPublisher) Write(o oops.Oops) error {
	p.Storage[o.Id] = &o
	// Returning error nil for interface compliance
	return nil
}

// Read reads the binary marshalling from the disk.
func (p InMemoryPublisher) Read(id string) (*oops.Oops, error) {
	o, ok := p.Storage[id]
	if !ok {
		return nil, errors.New("no oops with this id in the storage")
	}
	return o, nil
}

func NewInMemoryPublisher() InMemoryPublisher {
	return InMemoryPublisher{make(InMemoryOopsStorage)}
}
