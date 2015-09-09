package oops_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/vds/oops"
)

func TestOopsMarshalAndUnmarshal(t *testing.T) {
	err := errors.New("this is an error")
	o0 := oops.Oops{
		Error:          err.Error(),
		Panic:          true,
		Stack:          "stack",
		RequestDetails: map[string]string{"foo": "bar"},
	}

	encoded, err := o0.Marshal()
	if err != nil {
		t.Fatalf("failed to marshal: %s", err)
	}
	var o1 oops.Oops
	err = o1.Unmarshal(encoded)
	if err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}
	if !reflect.DeepEqual(o0, o1) {
		t.Errorf("unmarshalled OOPS does not match original one.")
	}
}
