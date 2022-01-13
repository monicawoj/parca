package columnstore

type Granule struct {
	Schema *Schema
	Parts  []*Part

	totalRows int
}

func NewGranule(s *Schema) *Granule {
	return &Granule{
		Schema: s,
	}
}

type Part struct {
	Schema  *Schema
	Columns []Column
}

func NewPart(s *Schema) *Part {
	p := &Part{
		Schema:  s,
		Columns: make([]Column, 0, len(s.ColumnDefinitions)),
	}

	for _, c := range s.ColumnDefinitions {
		p.Columns = append(p.Columns, c.NewColumn())
	}

	return p
}

func (p *Part) Insert(rows ...Row) error {
	// TODO

	return nil
}

// Insert inserts rows as an atomic part. It returns the number of rows total
// and an error if unsuccessful.
func (g *Granule) Insert(rows ...Row) (int, error) {
	p := NewPart(g.Schema)
	err := p.Insert(rows...)
	if err != nil {
		return g.totalRows, err
	}
	g.Parts = append(g.Parts, p)

	g.totalRows += len(rows)
	return g.totalRows, nil
}

// Compact merges all parts into a single, sorted part.
func (g *Granule) Compact() error {
	// TODO: Interate over all rows in all parts and add them to a single, new part.
	return nil
}

// TODO
//type GranuleIterator struct {
//	parts []*Part
//}
//
//// Iterator returns an iterator that will iterate over the rows of all parts of the granule in order of the ordered columns defined in the schema.
//func (g *Granule) Iterator() *GranuleIterator {
//
//}
