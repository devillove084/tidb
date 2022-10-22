package foreign

type Entry struct {
	schema ForeignSchema
	name   string
}

func (e *Entry) NewEntry(schema ForeignSchema, name string) *Entry {
	return &Entry{
		schema: schema,
		name:   name,
	}
}

func (e *Entry) Path() []string {
	return e.schema.Path(e.name)
}

type TableEntry struct {
	entry *Entry
	table Table
	sqls  []string
}

func (te *TableEntry) NewTableEntry(schema ForeignSchema, name string, sqls []string) *TableEntry {
	e := te.entry.NewEntry(schema, name)
	return &TableEntry{
		entry: e,
		sqls:  sqls,
		table: nil,
	}
}

func (te *TableEntry) GetTable() Table {
	return te.table
}

type TypeEntry struct {
	entry         *Entry
	protoDataType RelProtoDataType
}

func (tye *TypeEntry) NewTypeEntry(schema ForeignSchema, name string, protoDataType RelProtoDataType) *TypeEntry {
	e := tye.entry.NewEntry(schema, name)
	return &TypeEntry{
		entry:         e,
		protoDataType: protoDataType,
	}
}

func (tye *TypeEntry) GetType() *RelProtoDataType {
	return &tye.protoDataType
}
