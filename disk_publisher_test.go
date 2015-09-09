package oops_test

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/vds/oops"
)

func TestDiskPublisherWrite(t *testing.T) {
	o := oops.Oops{
		Id:    "oopsId",
		Error: errors.New("this is an error").Error(),
	}
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
	dp := oops.DiskPublisher{tempFolder}
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

func TestDiskPublisherRead(t *testing.T) {
	o0 := oops.Oops{
		Id:    "oopsId",
		Error: errors.New("this is an error").Error(),
	}
	tempFolder, err := ioutil.TempDir("/tmp", "oops")
	if err != nil {
		t.Error("error creating temporary directory for oopses")
	}
	defer os.RemoveAll(tempFolder)
	dp := oops.DiskPublisher{tempFolder}
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
