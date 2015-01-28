package publishers

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/vds/oops"
)

// DiskPublisher is a very simple publisher that is able to read and write the oops binary marshalling to and from the disk.
type DiskPublisher struct {
	OopsFolder string
}

// Write writes the binary marshalling of a oops on the disk.
func (p DiskPublisher) Write(o oops.Oops) error {
	data, err := o.Marshal()
	if err != nil {
		return fmt.Errorf("cannot marshal oops: %s", err)
	}
	oopsPath := path.Join(p.OopsFolder, o.Id)
	err = ioutil.WriteFile(oopsPath, data, 0600)
	if err != nil {
		return fmt.Errorf("cannot write oops to disk: %s", err)
	}
	return nil
}

// Read reads the binary marshalling from the disk.
func (p DiskPublisher) Read(id string) (*oops.Oops, error) {
	o := oops.Oops{}
	oopsPath := path.Join(p.OopsFolder, id)
	encoded_oops, err := ioutil.ReadFile(oopsPath)
	if err != nil {
		return nil, err
	}
	err = o.Unmarshal(encoded_oops)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func NewDiskPublisher(oopsFolder string) DiskPublisher {
	return DiskPublisher{OopsFolder: oopsFolder}
}
