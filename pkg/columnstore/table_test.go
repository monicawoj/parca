package columnstore

import (
	"fmt"
	"testing"

	"github.com/parca-dev/parca/pkg/columnstore/encoding"
	"github.com/parca-dev/parca/pkg/columnstore/types"
	"github.com/stretchr/testify/require"
)

func TestTable(t *testing.T) {
	labelsColumn := NewColumnDef("labels", NewMapColumnType(types.String, types.String, encoding.DictionaryRLEEncoding))
	timestampColumn := NewColumnDef("timestamp", NewPlainColumnType(types.Uint64, encoding.PlainEncoding))
	valueColumn := NewColumnDef("value", NewPlainColumnType(types.Int64, encoding.PlainEncoding))

	schema := NewSchema(
		ColumnDefinitions{
			labelsColumn,
			timestampColumn,
			valueColumn,
		},
		WithGranuleSize(2^13), // 8192
		WithOrderedColumns(
			labelsColumn,
			timestampColumn,
		),
	)

	table := NewTable("test", schema)

	err := table.Insert(
		schema.NewRow(map[string]string{"label1": "value1", "label2": "value2"}, uint64(1), int64(1)),
		schema.NewRow(map[string]string{"label1": "value1", "label2": "value2", "label3": "value3"}, uint64(2), int64(2)),
	)
	require.NoError(t, err)

	fmt.Println(table)
}
