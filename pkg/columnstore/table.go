package columnstore

type Table struct {
	Name     string
	Schema   *Schema
	Granules []*Granule
}

func NewTable(name string, schema *Schema) *Table {
	return &Table{
		Name:     name,
		Schema:   schema,
		Granules: []*Granule{},
	}
}

// Insert inserts rows into the table. Rows are expected to already be sorted
// by the table schema's sort columns.
func (t *Table) Insert(rows ...Row) error {
	// Ensure that dynamic column's schema is a superset of the schema of
	// `rows`.

	// Split rows into batches grouped respectively by the granule they belong
	// to and insert them into that granule.

	g := NewGranule(t.Schema)
	_, err := g.Insert(rows...)
	if err != nil {
		return err
	}
	// TODO: If granule is full, compact and split.
	t.Granules = append(t.Granules, g)

	return nil
}
