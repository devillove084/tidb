package foreign

type Comparable[V any] interface {
	Compare(b any) int
}

type Pair[T, U any] struct {
	First  T
	Second U
}
