package oops

import (
	"fmt"
	"os"
	"time"
)

// Generate a pseudo UUID used as id for oopses.
func _id() (id string) {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		panic("cannot open /dev/urandom")
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	return fmt.Sprintf("OOPS-%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// Factory is the component that holds the configuration and that is used to create the oopses.
type Factory struct {
	Publisher Publisher
}

// newOops creates a new oops from an error or a panic and delegate the publisher to persist
// said oops.
func (of Factory) newOops(e error, panic bool, requestDetails map[string]string) string {
	o := Oops{
		Id:             _id(),
		Time:           time.Now(),
		RequestDetails: requestDetails,
	}
	o.SetError(e, panic)
	of.Publisher.Write(o)
	return o.Id
}

// New creates a new oops from an error and delegates the publisher to persist
// it.
func (of Factory) New(e error, requestDetails map[string]string) string {
	return of.newOops(e, false, requestDetails)
}

// NewPanic creates a new oops for a panic and delegates the publisher to
// persist it.
func (of Factory) NewPanic(e error, requestDetails map[string]string) string {
	return of.newOops(e, true, requestDetails)
}
