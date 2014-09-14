package oops_test

import (
	"bytes"
	"encoding/gob"
	"errors"
	"reflect"
	"testing"

	"github.com/vds/oops"
)

// Tests the creation of a oops and the recording of information from an error from a panic.
func TestOopsPutError(t *testing.T) {

	errorString := "this is an error"
	err := errors.New(errorString)
	o := oops.Oops{}
	o.SetError(err, true)

	if o.Stack == "" {
		t.Error("Stack is empty.")
	}

	if o.Error != errorString {
		t.Errorf("Wrong Error: %v.\n", o.Error)
	}

	if o.ErrorType == "" {
		t.Error("ErrorType is empty.")
	}

	if o.Panic == false {
		t.Error("ErrorType is empty.")
	}
}

// Tests the binary marshalling of a oops.
func TestOopsMarshal(t *testing.T) {

	err := errors.New("this is an error")
	o0 := oops.Oops{}
	o0.SetError(err, true)

	encoded_oops, err := o0.Marshal()
	if err != nil {
		t.Error(err)
	}
	dec := gob.NewDecoder(bytes.NewReader(encoded_oops))
	var o1 oops.Oops
	err = dec.Decode(&o1)
	if err != nil {
		t.Errorf("Failed to decode: %s\n", err)
	}
	if !reflect.DeepEqual(o0, o1) {
		t.Errorf("Decoding does not match.")
	}
}

// Tests the binary unmarshalling of a oops.
func TestOopsUnmarshal(t *testing.T) {

	err := errors.New("this is an error")
	o0 := oops.Oops{}
	o0.SetError(err, true)

	encoded_oops, err := o0.Marshal()
	if err != nil {
		t.Error(err)
	}
	var o1 oops.Oops
	err = o1.Unmarshal(encoded_oops)
	if err != nil {
		t.Errorf("Failed to decode: %s\n", err)
	}

	if !reflect.DeepEqual(o0, o1) {
		t.Errorf("Decoding does not match.")
	}
}
