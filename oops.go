package oops

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"runtime"
	"time"
)

// Oops collects information about an error or a panic.
type Oops struct {
	Id             string
	RequestDetails map[string]string
	Time           time.Time
	Stack          string
	Error          string
	ErrorType      string
	Panic          bool
}

// SetError records the information about the error or the panic.
func (o *Oops) SetError(e error, panic bool) {
	o.Panic = panic
	o.Error = e.Error()
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
	o.ErrorType = reflect.TypeOf(e).String()
}

// Marshal returns a gob encoding of the oops.
func (o Oops) Marshal() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(o)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

// Unmarshal returns a gob decoding of the oops.
func (o *Oops) Unmarshal(encoded_oops []byte) (err error) {
	dec := gob.NewDecoder(bytes.NewReader(encoded_oops))
	err = dec.Decode(o)
	if err != nil {
		return err
	}
	return nil
}
