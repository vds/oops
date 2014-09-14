package factory

import (
	"fmt"
	"os"

	"github.com/vds/oops"
	"github.com/vds/oops/publishers"
)

// Generate a pseudo UUID used as id for oopses.
func _id() (id string) {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		panic("cannot open /dev/random")
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	return fmt.Sprintf("OOPS-%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// OopsFactory is the component that holds the configuration and that is used to create the oopses.
type OopsFactory struct {
	Publisher publishers.Publisher
}

// NewOops creates a new oops from and error or a panic and delegate the publisher to persiste said oops.
func (of OopsFactory) NewOops(e error, panic bool) (id string) {
	o := oops.Oops{}
	o.Id = _id()
	o.SetError(e, panic)
	of.Publisher.Write(o)
	return o.Id
}
