package encoding

import (
	"github.com/parca-dev/parca/pkg/columnstore/types"
)

type Encoding uint64

const (
	PlainEncoding Encoding = iota
	RLEEncoding
	DictionaryEncoding
	DictionaryRLEEncoding
)

// Array is the abstraction of the encoding array
type Array interface {
	Insert(int, types.Value) (int, error)
	Find(types.Value) (IndexRange, error)
}
