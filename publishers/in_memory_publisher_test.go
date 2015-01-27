package publishers_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/vds/oops"
	"github.com/vds/oops/publishers"
)

//Tests that the InMemoryOopsStorage can store an oops correctly.
func TestInMemoryPublisherWrite(t *testing.T) {
	e := errors.New("this is an error")
	o := oops.Oops{}
	o.SetError(e, true)
	o.Id = "oopsId"
	s := make(publishers.InMemoryOopsStorage)
	p := publishers.InMemoryPublisher{s}
	p.Write(o)
	if !reflect.DeepEqual(o, *s[o.Id]) {
		t.Error("oops not stored correctly")
	}
}

//Tests that the InMemoryOopsStorage can read an oops correctly.
func TestInMemoryPublisherRead(t *testing.T) {
	e := errors.New("this is an error")
	o := oops.Oops{}
	o.SetError(e, true)
	o.Id = "oopsId"
	s := make(publishers.InMemoryOopsStorage)
	p := publishers.InMemoryPublisher{s}
	p.Write(o)
	newOops, err := p.Read(o.Id)
	if err != nil {
		t.Error("error reading the oops: %#v", err)
	}
	if !reflect.DeepEqual(*newOops, o) {
		t.Error("oops not read correctly")
	}
}
