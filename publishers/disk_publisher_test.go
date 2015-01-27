package publishers_test

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/vds/oops"
	"github.com/vds/oops/publishers"
)

//Tests that the DiskPublisher can write a oops on the disk correctly.
func TestDiskPublisherWrite(t *testing.T) {
	e := errors.New("this is an error")
	o := oops.Oops{}
	o.SetError(e, true)
	o.Id = "oopsId"
	tempFolder, err := ioutil.TempDir("/tmp", "oops")
	if err != nil {
		t.Error("error creating temporary directory for oopses")
	}
	defer os.RemoveAll(tempFolder)

	fileInfo, err := ioutil.ReadDir(tempFolder)
	if err != nil {
		t.Error("error reading temporary directory")
	}
	if len(fileInfo) != 0 {
		t.Error("error, temporary directory not empty")
	}
	dp := publishers.DiskPublisher{tempFolder}
	err = dp.Write(o)
	if err != nil {
		t.Error("error writing the oops")
	}
	fileInfo, err = ioutil.ReadDir(tempFolder)
	if err != nil {
		t.Error("error reading temporary directory")
	}
	if len(fileInfo) != 1 {
		t.Error("error, temporary directory should contain only one oops")
	}
}

//Tests that the DiskPublisher can read a oops from the disk correctly.
func TestDiskPublisherRead(t *testing.T) {
	e := errors.New("this is an error")
	o0 := oops.Oops{}
	o0.SetError(e, true)
	o0.Id = "oopsId"
	tempFolder, err := ioutil.TempDir("/tmp", "oops")
	if err != nil {
		t.Error("error creating temporary directory for oopses")
	}
	defer os.RemoveAll(tempFolder)
	dp := publishers.DiskPublisher{tempFolder}
	err = dp.Write(o0)
	if err != nil {
		t.Error("error writing oops")
	}
	o1, err := dp.Read(o0.Id)
	if err != nil {
		t.Error("error read oops")
	}
	if !reflect.DeepEqual(o0, *o1) {
		t.Errorf("decoding does not match")
	}
}
