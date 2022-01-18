package encoding

import (
	"testing"

	"github.com/parca-dev/parca/pkg/columnstore/types"
	"github.com/stretchr/testify/require"
)

func TestPlain(t *testing.T) {
	p := NewPlain(types.String)

	count, err := p.Insert(0, types.Value{Data: "test"})
	require.NoError(t, err)
	require.Equal(t, 1, count)
	require.Equal(t, []types.Value{
		{Data: "test"},
	}, p.values)
}

func TestPlainInsertTwo(t *testing.T) {
	p := NewPlain(types.String)

	count, err := p.Insert(0, types.Value{Data: "test1"})
	require.NoError(t, err)
	require.Equal(t, 1, count)

	count, err = p.Insert(1, types.Value{Data: "test3"})
	require.NoError(t, err)
	require.Equal(t, 2, count)

	require.Equal(t, []types.Value{
		{Data: "test1"},
		{Data: "test3"},
	}, p.values)
}

func TestPlainSparseInsert(t *testing.T) {
	typ := types.String
	p := NewPlain(typ)

	count, err := p.Insert(2, types.Value{Data: "test3"})
	require.NoError(t, err)
	require.Equal(t, 3, count)

	require.Equal(t, []types.Value{
		typ.Null(),
		typ.Null(),
		{Data: "test3"},
	}, p.values)
}
