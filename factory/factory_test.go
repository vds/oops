package factory_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/vds/oops/factory"
	"github.com/vds/oops/publishers"
)

// Tests the whole workflow, from the creation of the Publisher, the OopsFactory, the creation of a
// oops and the recording of information from an error from a panic.
func TestCreatioAndPublicationOfAOops(t *testing.T) {
	errorString := "this is an error"
	e := errors.New(errorString)
	tempFolder, err := ioutil.TempDir("/tmp", "oops")
	if err != nil {
		t.Error("Error creating temporary directory for oopses.")
	}
	defer os.RemoveAll(tempFolder)
	p := publishers.DiskPublisher{tempFolder}
	of := factory.OopsFactory{p}
	id := of.NewOops(e, true)
	matched, err := regexp.MatchString(`^OOPS-\S{8}-\S{4}-\S{4}-\S{4}-\S{12}$`, id)
	if err != nil {
		t.Error("Error creating regexp.")
	}
	if !matched {
		t.Error("Error matching regexp.")
	}
	o, err := p.Read(path.Join(tempFolder, id))
	if o.Error != errorString {
		t.Error("Error matching error string.")
	}
	if o.Stack == "" {
		t.Error("Error, empty stack.")
	}
}