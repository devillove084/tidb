package foreign

type NameMultiMap[V any] struct {
	nmmap NameMap[[]V]
}
