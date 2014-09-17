package publishers

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/vds/oops"
)

// Publisher is the componet in charge of persisting the oops.
type Publisher interface {
	Write(o oops.Oops)
}

// DiskPublisher is a very simple publisher that is able to read and write the oops binary marshalling to and from the disk.
type DiskPublisher struct {
	OopsFolder string
}

// Write writes the binary marshalling of a oops on the disk.
func (p DiskPublisher) Write(o oops.Oops) {
	data, err := o.Marshal()
	if err != nil {
		panic(fmt.Sprintf("cannot write oops to disk: %s", err))
	}
	oopsPath := path.Join(p.OopsFolder, o.Id)
	ioutil.WriteFile(oopsPath, data, 0600)
}

// Read reads the binary marshalling from the disk.
func (p DiskPublisher) Read(path string) (o oops.Oops, err error) {
	encoded_oops, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = o.Unmarshal(encoded_oops)
	return
}
