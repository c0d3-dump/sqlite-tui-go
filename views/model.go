package views

type focusedView int

type Model struct {
	cursor int
	tables []Table
}

type Table struct {
	cursor  int
	columns []string
	ids     []string
	data    TableData
}

type TableData struct {
	cursor int
	data   map[string]string
}

func NewModel() Model {
	return Model{
		cursor: -1,
		tables: []Table{},
	}
}

type CreateTable struct {
	name    string
	columns []CreateColumn
}

type CreateColumn struct {
	name    string
	dtype   string
	notnull bool
	dval    any
	pk      bool
}

func newCreateTable() CreateTable {
	return CreateTable{
		name:    "",
		columns: []CreateColumn{},
	}
}
