package foreign

type LatticeEntry struct {
	entry *Entry
}

func (le *LatticeEntry) GetLattice() *Lattice {
	return nil
}

func (le *LatticeEntry) GetStarTable() *TableEntry {
	return nil
}

type ForeignSchema struct {
	parent             *ForeignSchema
	schema             Schema
	name               string
	tableMap           NameMap[TableEntry]
	functionMap        NameMultiMap[FunctionEntry]
	typeMap            NameMap[TypeEntry]
	latticeMap         NameMap[LatticeEntry]
	functionNames      NameSet
	nullaryFunctionMap NameMap[FunctionEntry]
	subSchemaMap       NameMap[ForeignSchema]
	path               []string
}

func (fs *ForeignSchema) Path(name string) []string {
	// TODO:
	return nil
}
