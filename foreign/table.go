package foreign

type Table interface {
	GetRowType(typeFactory RelDataTypeFactory) (RelDataType, error)
	// GetStatistic() Statistic
	GetGdbcTableType() (TableType, error)
	IsRolledUp(column string) bool

	RolledUpColumnValidInsideAgg(column string, call SqlCall, parent SqlNode, config ForeignConnectionConfig)
}
