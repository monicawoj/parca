package columnstore

import (
	"fmt"

	"github.com/parca-dev/parca/pkg/columnstore/encoding"
	"github.com/parca-dev/parca/pkg/columnstore/types"
)

// Column ...
type Column interface {
	InsertAt(index int, value types.Value) error
}

// PlainColumn ...
type PlainColumn struct {
	typ *PlainColumnType
	enc encoding.Array
}

// NewPlainColumn ...
func NewPlainColumn(ctyp *PlainColumnType) *PlainColumn {
	return &PlainColumn{
		typ: ctyp,
		enc: encoding.NewPlain(ctyp.typ, 10), // TODO arbitrary number is arbitrary
	}
}

// InsertAt ...
func (p *PlainColumn) InsertAt(index int, value types.Value) error {
	_, err := p.enc.Insert(index, value)
	return err
}

// MapColumn ...
type MapColumn struct {
	typ     *MapColumnType
	columns map[interface{}]encoding.Array
}

// NewMapColumn ...
func NewMapColumn(typ *MapColumnType) *MapColumn {
	return &MapColumn{
		typ:     typ,
		columns: map[interface{}]encoding.Array{},
	}
}

// InsertAt ...
func (m *MapColumn) InsertAt(index int, value types.Value) error {

	mt, ok := value.Type.(*types.MapType)
	if !ok {
		return fmt.Errorf("value was not of type MapType")
	}

	switch mt.Key {
	case types.String:
		switch mt.Value {
		case types.String:
			v, ok := value.Data.(map[string]string)
			if !ok {
				return fmt.Errorf("unknown map")
			}

			// Insert values into all key columns
			for key, val := range v {
				array, ok := m.columns[key]
				if !ok {
					array = encoding.NewDictionaryRLE(types.String, -1)
					m.columns[key] = array
				}

				_, err := array.Insert(index, types.Value{Data: val})
				if err != nil {
					return fmt.Errorf("inset failed: %w", err)
				}
			}

		default:
			panic(fmt.Sprintf("unsupported map type: %v", mt.Value))
		}
	default:
		panic(fmt.Sprintf("unsupported map type: %v", mt.Key))
	}

	return nil
}
