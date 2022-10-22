package foreign

import (
	"github.com/modern-go/reflect2"
	"github.com/pingcap/tidb/expression"
)

type Enumerator[T any] interface {
	Current() (T, error)
	MoveNext() (T, error)
	Reset()
	Close()
}

type RawEnumerable[T any] interface {
	Enumerator() (Enumerator[T], error)
}

type ExtendedEnumerable[T any] interface {
}

type ExtendedQueryable[T any] interface {
}

type Enumerable[T any] interface {
	RawEnumerable[T]
	// Iterable[T]
	ExtendedEnumerable[T]
	AsQueryable() (Queryable[T], error)
}

type RawQueryable[T any] interface {
	Enumerable[T]
	GetElementType() (reflect2.Type, error)
	GetExpression() (expression.Expression, error)
	GetProvider() (QueryProvider[T], error)
}

type Queryable[T any] interface {
	RawQueryable[T]
	ExtendedQueryable[T]
}

type QueryProvider[T any] interface {
	CreateQueryByReflect(expression expression.Expression, rowType reflect2.Type) (Queryable[T], error)
	CreateQueryByStruct(expression expression.Expression, rowType interface{}) (Queryable[T], error)
	ExecuteInReflect(expression expression.Expression, t reflect2.Type) (T, error)
	ExecuteInStruct(expression expression.Expression, t interface{}) (T, error)
	ExecuteQuery(qa *Queryable[T]) (Enumerator[T], error)
}
