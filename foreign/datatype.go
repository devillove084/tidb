package foreign

import "github.com/pingcap/tidb/expression"

type TableType int

const (
	TableT TableType = iota
	VIEW
	FOREIGEN_TABLE
	MATERIALIZED_VIEW
	INDEX
	JOIN
	SEQUENCE
	STAR
	STREAM
	TYPE
	SYSTEM_TABLE
	SYSTEM_VIEW
	SYSTEM_INDEX
	SYSTEM_TOAST_INDEX
	SYSTEM_TOAST_TABLE
	TEMPORARY_INDEX
	TEMPORARY_SEQUENCE
	TEMPORARY_TABLE
	TEMPORARY_VIEW
	LOCAL_TEMPORARY
	SYNONYM
	ALIAS
	GLOBAL_TEMPORARY
	ACCEL_ONLY_TABLE
	AUXILIARY_TABLE
	GLOBAL_TEMPORARY_TABLE
	HIERARCHY_TABLE
	INOPERATIVE_VIEW
	MATERIALIZED_QUERY_TABLE
	NICKNAME
	TYPED_TABLE
	TYPED_VIEW
	TEMPORAL_TABLE
	OTHER
)

type RelDataTypeComparability string

const (
	NONEC     RelDataTypeComparability = "No comparisons allowed"
	UNORDERED RelDataTypeComparability = "Only equals/not-equals allowed"
	ALL       RelDataTypeComparability = "All comparisons allowed"
)

// func (rdtc *RelDataTypeComparability) NewRelDataTypeComparability(desctiption string) {

// }

type RelDataTypeSystem interface {
}

type ToFidldName struct {
	function Function1[RelDataTypeField, string]
}

func (tof *ToFidldName) Apply(o RelDataTypeField) string {
	return o.GetName()
}

type MapEntry[K, V any] interface {
	GetKey() K

	GetValue() V

	SetValue(value V) V

	Equals(o interface{}) bool

	HashCode() int

	Comparator() bool
}

type SqlCollation struct{}

type SqlIntervalQualifier struct{}

type SqlIdentifier struct{}

type RelDataTypeFamily interface{}

type RelDataTypePrecedenceList interface {
	ContainsType(t RelDataType) bool
	CompareTypePrecedence(t1, t2 RelDataType)
}

type RelDataType interface {
	IsStruct() bool

	GetFieldList() []RelDataTypeField

	GetFieldNames() []string

	GetFieldCount() int

	GetStructKind() (StructKind, error)

	GetField(fieldName string, caseSensitive bool, elideRecord bool) (RelDataTypeField, error)

	IsNullable() bool

	GetComponentType() (RelDataType, error)

	GetKeyType() (RelDataType, error)

	GetValueType() (RelDataType, error)

	GetCharset() (byte, error)

	GetCollation() (SqlCollation, error)

	GetIntervalQualifier() (SqlIntervalQualifier, error)

	GetPrecision() int

	GetScale() int

	// GetSqlTypeName() (SqlTypeName, error)

	GetSqlIdentifier() (SqlIdentifier, error)

	ToString() (string, error)

	GetFullTypeString() (string, error)

	GetFamily() (RelDataTypeFamily, error)

	GetPrecedenceList() (RelDataTypePrecedenceList, error)

	GetComparability() (RelDataTypeComparability, error)

	IsDynamicStruct() bool

	EqualsSansFieldNames(that RelDataType) bool
}

type RelOptTable interface {
	Wrapper
	GetQualifiedName() []string
	GetRowCount() float64
	GetRowType() (RelDataType, error)
	GetRelOptSchema() (RelDataType, error)
}

type RelNode interface {
	RelOptTable
}

type RelDataTypeField interface {
	//map[string]RelDataType
	MapEntry[string, RelDataType]

	GetName() string

	GetIndex() int

	GetType() RelDataType

	IsDynamicStar() bool
}

type Function[R any] interface {
	// TODO: a function root type
	ApplyF(o R) string
}

type Function1[T any, R any] interface {
	Function[R]
	Identity() (Function1[T, R], error)
	Apply(a T) (R, error)
}

type RelProtoDataType interface {
	Function1[RelDataTypeFactory, RelDataType]
}

type SchemaPlus interface {
	Schema
}

type SchemaVersion interface {
	IsBefore(other SchemaVersion) bool
}

type SqlNodeImpl interface {
}

type SqlNode struct {
}

type SqlCallImpl interface {
}

type SqlCall struct {
}

type FunctionParameter interface {
}

type FunctionEntry interface {
	GetParameters() ([]FunctionParameter, error)
}

type RelRunner interface {
	//PrepareStatement(rel RelNode) (PreparedStatement, error)
}

type StructKind = int

const (
	NONE StructKind = iota
	FULLY_QUALIFIED
	PEEK_FIELDS_DEFAULT
	PEEK_FIELDS
	PEEK_FIELDS_NO_EXPAND
)

// RelDataTypeFactory is a factory for datatype descriptors. It defines methods
// for instantiating and combining SQL, Java, and collection types. The factory
// also provides methods for return type inference for arithmetic in cases where
// SQL 2003 is implementation defined or impractical.
//
// This interface is an example of the
// {Glossary#ABSTRACT_FACTORY_PATTERN abstract factory pattern}.
// Any implementation of `RelDataTypeFactory` must ensure that type
// objects are canonical: two types are equal if and only if they are
// represented by the same Java object. This reduces memory consumption and
// comparison cost.
type RelDataTypeFactory interface {
	GetTypeSystem() (RelDataTypeSystem, error)
	CreateGoType(clazz interface{}) (RelDataType, error)
	CreateJoinType(types ...RelDataType) (RelDataType, error)
	CreateStructTypeByKind(kind StructKind, typeList []RelDataType, fieldNameList []string) (RelDataType, error)
	CreateStructTypeWithoutKind(typeList []RelDataType, fieldNameList []string) (RelDataType, error)
	CreateStructTypeByFieldList([]map[string]RelDataType) (RelDataType, error)
	CreateArrayType(elementType RelDataType, maxCardinality uint32) (RelDataType, error)
	CreateMapType(keyType RelDataType, valueType RelDataType) (RelDataType, error)
	CreateMultiSetType(elementType RelDataType, maxCardinality uint32) (RelDataType, error)
	CopyType(t RelDataType) (RelDataType, error)
}

// The typical way for a table to be created is when TiDB interrgogates a
// user-defined schema in order to validate names appearing in a SQL query.
//
// TiDB finds the schema by calling `GetSubSchema(string)` on the connection's
// root schema, then gets a table by calling `GetTable(string)`
//
// Note that a table does not know its name. It is in fact possible for a
// table to be used more than once, perhaps under multiple names or under
// multiple schemas. (Compare with the i-node concept in the UNIX filesystem.)
//
// A particulat table instance may also implement `Wrapper` to give access to
// sub-objects.
type TableImpl interface {
	GetRowType(typeFactory RelDataTypeFactory) (RelDataType, error)
	// GetStatistic() (Statistic, error)
	GetGdbcTableType() (TableType, error)
	IsRolledUp(column string) bool
	RolledUpColumnValidInsideAgg(column string, call SqlCall, parent SqlNode, config ForeignConnectionConfig) bool
}

// A namespace for tables and functions.
//
// A schema can also contain sub-schemas, to any level of nesting. Most
// providers have limited number of levels; for example, most ORM framework
// databases have either one level ("schemas") or two levels ("databases" and "catalog")
//
// There may be multiple overloaded functions with the same name but
// different numbers or parameters.
// For this reason, `GetFunctions` returns a list of all
// members with the same name.
//
// The most common and important type of member is the one with no
// arguments and a result type that is a collection of records(recordset). This is called a `relation`.
// It is equivalent to table in a relational database.
//
// For example, the query `select * from sales.emps` is valid if sales is registerd schema and `emps` is a member
// with zero parameters and a result type of `Collection(Record(int: "empno", String: "name"))`
//
// By the way, a schema may be nested within another schema, in `GetSubSchema`
type Schema interface {
	GetTable(name string) (TableImpl, error)
	GetTableNames(name string) ([]string, error)
	GetType(name string) (RelProtoDataType, error)
	GetTypeNames() ([]string, error)
	GetFunctions(name string) ([]FunctionEntry, error)
	GetFunctionNames() ([]string, error)
	GetSubSchema(name string) (Schema, error)
	GetExpression(parentSchema SchemaPlus, name string) (expression.Expression, error)
	IsMutable() bool
	Snapshot(version SchemaVersion) (Schema, error)
}
