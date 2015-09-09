package oops_test

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/vds/oops"
)

// Tests the whole workflow, the creation of the Publisher, the Factory, the creation of a
// oops and the recording of information from an error from a panic.
func TestCreationAndPublicationOfAOops(t *testing.T) {
	errorString := "this is an error"
	e := errors.New(errorString)
	tempFolder, err := ioutil.TempDir("/tmp", "oops")
	if err != nil {
		t.Error("error creating temporary directory for oopses")
	}
	defer os.RemoveAll(tempFolder)
	p := oops.DiskPublisher{tempFolder}
	of := oops.Factory{p}
	reqDetails := map[string]string{"foo": "bar"}
	stack := "stack"
	id, err := of.New(e, oops.SetStack(stack), oops.SetRequestDetails(reqDetails))
	if err != nil {
		t.Fatalf("Failed to create new OOPS: %v", err)
	}
	matched, err := regexp.MatchString(`^OOPS-\S{8}-\S{4}-\S{4}-\S{4}-\S{12}$`, id)
	if err != nil {
		t.Error("error creating regexp")
	}
	if !matched {
		t.Error("error matching regexp")
	}
	o, err := p.Read(id)
	if o.Error != errorString {
		t.Error("error matching error string")
	}
	if o.Stack != stack {
		t.Errorf("Unexpected stack; got %s, want %s", o.Stack, stack)
	}
	if !reflect.DeepEqual(o.RequestDetails, reqDetails) {
		t.Errorf("Unexpected request details; got %s, want %s", o.RequestDetails, reqDetails)
	}
}
