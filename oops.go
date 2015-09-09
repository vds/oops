package oops

import (
	"bytes"
	"encoding/gob"
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
