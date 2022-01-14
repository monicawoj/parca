package columnstore

import (
	"github.com/parca-dev/parca/pkg/columnstore/encoding"
	"github.com/parca-dev/parca/pkg/columnstore/types"
)

type Column interface {
	InsertAt(index int, value types.Value) error
}

type PlainColumn struct {
	typ *PlainColumnType
	enc encoding.Array
}

func NewPlainColumn(ctyp *PlainColumnType) *PlainColumn {
	return &PlainColumn{
		typ: ctyp,
		enc: encoding.NewPlain(ctyp.typ, 10), // TODO arbitrary number is arbitrary
	}
}

func (p *PlainColumn) InsertAt(index int, value types.Value) error {
	_, err := p.enc.Insert(index, value)
	return err
}

type MapColumn struct {
	typ *MapColumnType
}

func (m *MapColumn) InsertAt(index int, value types.Value) error {
	return nil
}
