package columnstore

import (
	"github.com/parca-dev/parca/pkg/columnstore/encoding"
	"github.com/parca-dev/parca/pkg/columnstore/types"
)

type Schema struct {
	ColumnDefinitions ColumnDefinitions
	OrderedColumns    ColumnDefinitions
}

func NewSchema(
	columns ColumnDefinitions,
	options ...func(s *Schema),
) *Schema {
	s := &Schema{
		ColumnDefinitions: columns,
	}

	for _, option := range options {
		option(s)
	}

	return s
}

func WithOrderedColumns(
	columns ...*ColumnDefinition,
) func(s *Schema) {
	return func(s *Schema) {
		s.OrderedColumns = ColumnDefinitions(columns)
	}
}

func (s *Schema) NewRow(columnData ...interface{}) Row {
	r := Row{
		ColumnData: make([]types.Value, 0, len(columnData)),
	}

	for i, data := range columnData {
		r.ColumnData = append(r.ColumnData, s.ColumnDefinitions[i].ColumnType.ValueType().NewValue(data))
	}

	return r
}

type ColumnDefinitions []*ColumnDefinition

type ColumnDefinition struct {
	Name       string
	ColumnType ColumnType
}

func NewColumnDef(
	name string,
	columnType ColumnType,
) *ColumnDefinition {
	return &ColumnDefinition{
		Name:       name,
		ColumnType: columnType,
	}
}

func (c *ColumnDefinition) NewColumn() Column {
	return c.ColumnType.NewColumn()
}

type ColumnType interface {
	NewColumn() Column
	ValueType() types.Type
}

func NewPlainColumnType(
	t types.Type,
	encoding encoding.Encoding,
) *PlainColumnType {
	return &PlainColumnType{
		typ:      t,
		Encoding: encoding,
	}
}

type PlainColumnType struct {
	Encoding encoding.Encoding
	typ      types.Type
}

func (p *PlainColumnType) NewColumn() Column {
	return &PlainColumn{
		typ: p,
	}
}

func (p *PlainColumnType) ValueType() types.Type {
	return p.typ
}

func NewMapColumnType(
	keyType types.Type,
	valueType types.Type,
	encoding encoding.Encoding,
) *MapColumnType {
	return &MapColumnType{
		typ:      types.Map(keyType, valueType),
		Encoding: encoding,
	}
}

type MapColumnType struct {
	Encoding encoding.Encoding
	typ      types.Type
}

func (m *MapColumnType) NewColumn() Column {
	return &MapColumn{
		typ: m,
	}
}

func (m *MapColumnType) ValueType() types.Type {
	return m.typ
}
