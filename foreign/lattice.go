package foreign

type LatticeTable struct {
	t     RelOptTable
	alias string
}

type LatticeNode struct {
	table    LatticeTable
	startCol int
	endCol   int
	alias    string
	children []LatticeChildNode
	digest   string
}

type LatticeRootNode struct {
	parentNode  LatticeNode
	descendants []LatticeNode
	paths       []Path
}

type LatticeChildNode struct {
	parent LatticeNode
	link   []IntPair
}

type Column interface {
	Comparable[Column]
	ToBitSet(columns []Column) (byte, error)
	HashCode() int
	Equals(obj interface{}) bool
}

type LatticeStatisticProvider interface {
}

type ForeignColumn struct {
	ordinal int
	alias   string
}

type Measure struct {
	//agg    SqlAggFunction
	distinct bool
	name     string
	args     []ForeignColumn
	digest   string
}

type Tile struct {
	measures   []Measure
	dimensions []ForeignColumn
	bitSet     byte
}

type Lattice struct {
	rootSchema         ForeignSchema
	rootNode           LatticeRootNode
	columns            []ForeignColumn
	auto               bool
	algorithm          bool
	algorithmMaxMillis uint64
	rowCountEstimate   float64
	defaultMeasures    []Measure
	tiles              []Tile
	columnUses         map[int]bool
	statisticProvider  LatticeStatisticProvider
}
