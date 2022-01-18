package encoding

import (
	"errors"
	"fmt"
	"strings"

	"github.com/parca-dev/parca/pkg/columnstore/types"
)

type Plain struct {
	typ    types.Type
	values []types.Value
}

func NewPlain(typ types.Type) *Plain {
	return &Plain{
		typ:    typ,
		values: make([]types.Value, 0, 10), // TODO arbitrary number is arbitrary, this should be optimized using a pool of plain encoding objects to re-use rather than pre-allocating.
	}
}

func (c *Plain) String() string {
	s := "[ "
	for i := 0; i < len(c.values); i++ {
		s += fmt.Sprint(c.values[i])
		s += ","
	}
	s = strings.TrimSuffix(s, ",")

	s += " ]"

	return s
}

func (c *Plain) Insert(index int, v types.Value) (int, error) {
	if index < 0 {
		return -1, errors.New("index out of range")
	}
	if index < len(c.values) {
		return -1, errors.New("cannot insert out of order")
	}

	if index > len(c.values) {
		// This could be further optimized by noting the first index where the
		// value is a non-null value as columns are expected to be very sparse,
		// but this decision should be backed by data.
		for i := len(c.values); i < index; i++ {
			c.values = append(c.values, c.typ.Null())
		}
	}

	c.values = append(c.values, v)
	return len(c.values), nil
}
