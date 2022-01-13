package columnstore

import "github.com/parca-dev/parca/pkg/columnstore/types"

type Column interface {
	InsertAt(index int, value types.Value) error
}

type PlainColumn struct {
	typ *PlainColumnType
}

func (p *PlainColumn) InsertAt(index int, value types.Value) error {
	return nil
}

type MapColumn struct {
	typ *MapColumnType
}

func (m *MapColumn) InsertAt(index int, value types.Value) error {
	return nil
}
