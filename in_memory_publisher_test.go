package oops_test

import (
	"errors"
	"reflect"
	"sync"
	"testing"

	"github.com/vds/oops"
)

//Tests that the InMemoryOopsStorage can store an oops correctly.
func TestInMemoryPublisherWrite(t *testing.T) {
	e := errors.New("this is an error")
	o := oops.Oops{}
	o.SetError(e, true)
	o.Id = "oopsId"
	s := make(oops.InMemoryOopsStorage)
	p := oops.InMemoryPublisher{s, sync.Mutex{}}
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
	s := make(oops.InMemoryOopsStorage)
	p := oops.InMemoryPublisher{s, sync.Mutex{}}
	p.Write(o)
	newOops, err := p.Read(o.Id)
	if err != nil {
		t.Error("error reading the oops: %#v", err)
	}
	if !reflect.DeepEqual(*newOops, o) {
		t.Error("oops not read correctly")
	}
}
