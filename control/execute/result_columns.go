package execute

type ResultColumns struct {
	Columns          []string
	GroupColumns     []string
	ResultColumns    []string
	DimensionColumns []string
	TagColumns       []string
}

func newResultColumns(groupColumns, resultColumns, dimensionColumns, tagColumns []string) *ResultColumns {
	columns := append([]string{}, groupColumns...)
	columns = append(columns, resultColumns...)
	columns = append(columns, dimensionColumns...)
	columns = append(columns, tagColumns...)

	return &ResultColumns{
		GroupColumns:     groupColumns,
		ResultColumns:    resultColumns,
		DimensionColumns: dimensionColumns,
		TagColumns:       tagColumns,
		Columns:          columns,
	}
}
