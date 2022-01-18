package encoding

import (
	"testing"

	"github.com/parca-dev/parca/pkg/columnstore/types"
	"github.com/stretchr/testify/require"
)

func Test_DictionaryRLE(t *testing.T) {
	p := NewDictionaryRLE(types.String)

	i := 0
	count, err := p.Insert(i, types.Value{Data: "test1"})
	i++
	require.NoError(t, err)
	require.Equal(t, i, count)

	for ; i < 4; i++ {
		count, err = p.Insert(i, types.Value{Data: "test2"})
		require.NoError(t, err)
		require.Equal(t, i+1, count)
	}

	count, err = p.Insert(4, types.Value{Data: "test3"})
	i++
	require.NoError(t, err)
	require.Equal(t, i, count)
}

func Test_DictionaryRLE_Insert(t *testing.T) {
	p := NewDictionaryRLE(types.String)

	i := 0
	count, err := p.Insert(i, types.Value{Data: "test1"})
	i++
	require.NoError(t, err)
	require.Equal(t, i, count)

	count, err = p.Insert(1, types.Value{Data: "test3"})
	i++
	require.NoError(t, err)
	require.Equal(t, i, count)

	count, err = p.Insert(1, types.Value{Data: "test2"})
	i++
	require.NoError(t, err)
	require.Equal(t, i, count)

	require.Equal(t, "test1,test2,test3,\n", p.String())
}

func Test_DictionaryRLE_AppendAt(t *testing.T) {
	p := NewDictionaryRLE(types.String)

	i := 1
	count, err := p.Insert(i, types.Value{Data: "test1"})
	require.NoError(t, err)
	require.Equal(t, i, count)
	require.Equal(t, "<nil>,test1,\n", p.String())
}
