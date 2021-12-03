package fizz

type SchemaQuery interface {
	ReplaceSchema(map[string]*Table)
	Build() error
	TableInfo(string) (*Table, error)
	ReplaceColumn(table string, oldColumn string, newColumn Column) error
	ColumnInfo(table string, column string) (*Column, error)
	IndexInfo(table string, idx string) (*Index, error)
	Delete(string)
	ResetCache()
	SetTable(*Table)
	DeleteColumn(string, string)
}
