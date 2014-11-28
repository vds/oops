package storage

import (
	"github.com/vds/oops"
)

type Storage interface {
	Read(id string) (oops.Oops, error)
	Write(oops.Oops) error
}
