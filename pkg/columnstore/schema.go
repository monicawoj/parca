package columnstore

import (
	"github.com/parca-dev/parca/pkg/columnstore/encoding"
	"github.com/parca-dev/parca/pkg/columnstore/types"
)

type Schema struct {
	ColumnDefinitions ColumnDefinitions
	OrderedColumns    ColumnDefinitions
	GranuleSize       int
}

func NewSchema(
	columns ColumnDefinitions,
	options ...func(s *Schema),
) *Schema {
	s := &Schema{
		ColumnDefinitions: columns,
		GranuleSize:       2 ^ 13,
	}

	for _, option := range options {
		option(s)
	}

	for _, columnDef := range columns {
		columnDef.Schema = s
	}

	return s
}

func WithGranuleSize(granuleSize int) func(s *Schema) {
	return func(s *Schema) {
		s.GranuleSize = granuleSize
	}
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
	Schema     *Schema
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
	return c.ColumnType.NewColumn(c)
}

type ColumnType interface {
	NewColumn(def *ColumnDefinition) Column
	ValueType() types.Type
}

func NewPlainColumnType(
	t types.Type,
	encoding encoding.Encoding,
) *PlainColumnType {
	return &PlainColumnType{
		typ:      t,
		encoding: encoding,
	}
}

type PlainColumnType struct {
	encoding encoding.Encoding
	typ      types.Type
}

func (p *PlainColumnType) NewColumn(def *ColumnDefinition) Column {
	return NewPlainColumn(def, p)
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
		encoding: encoding,
	}
}

type MapColumnType struct {
	encoding encoding.Encoding
	typ      types.Type
}

func (m *MapColumnType) NewColumn(def *ColumnDefinition) Column {
	return NewMapColumn(def, m)
}

func (m *MapColumnType) ValueType() types.Type {
	return m.typ
}
