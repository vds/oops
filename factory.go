package oops

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"time"
)

// Generate a pseudo UUID used as id for oopses.
func _id() (id string) {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		panic("cannot open /dev/urandom")
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	return fmt.Sprintf("OOPS-%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// Factory is the component that holds the configuration and that is used to create the oopses.
type Factory struct {
	Publisher Publisher
}

// new creates a new oops from an error or a panic and calls the publisher's
// Write method to persist it.
func (f Factory) new(e error, isPanic bool, options ...Option) (string, error) {
	o := &Oops{
		Id:   _id(),
		Time: time.Now(),
	}
	o.Error = e.Error()
	o.ErrorType = reflect.TypeOf(e).String()
	o.Panic = isPanic
	for _, option := range options {
		option(o)
	}
	if err := f.Publisher.Write(*o); err != nil {
		return "", nil
	}
	return o.Id, nil
}

// New creates a new oops from an error and delegates the publisher to persist
// it, returning the OOPS ID if nothing goes wrong.
func (f Factory) New(e error, options ...Option) (string, error) {
	return f.new(e, false, options...)
}

// NewPanic creates a new oops for a panic and delegates the publisher to
// persist it, returning the OOPS ID if nothing goes wrong.
func (f Factory) NewPanic(e error, options ...Option) (string, error) {
	return f.new(e, true, options...)
}

// Option is the signature of the optional functions passed to the Factory's
// New method.
type Option func(*Oops)

// SetStack returns an Option that assigns the given stack trace to the Oops.
func SetStack(stack string) Option {
	return func(o *Oops) {
		o.Stack = stack
	}
}

// SetRequestDetails returns an Option that assigns the given map of request
// details to the Oops.
func SetRequestDetails(m map[string]string) Option {
	return func(o *Oops) {
		o.RequestDetails = m
	}
}

// SetRuntimeStack is an Option that obtains stack traces from all goroutines
// and assign that to Oops.Stack.
func SetRuntimeStack(o *Oops) {
	stack := make([]byte, 1<<20)
	for i := 0; ; i++ {
		n := runtime.Stack(stack, true)
		if n < len(stack) {
			stack = stack[:n]
			break
		}
		if len(stack) >= 64<<20 {
			// Filled 64 MB - stop there.
			break
		}
		stack = make([]byte, 2*len(stack))
	}
	o.Stack = string(stack)
}
