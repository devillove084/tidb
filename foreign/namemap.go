package foreign

import (
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type NameMap[V any] struct {
	nvmap *orderedmap.OrderedMap[string, V]
}

func (nm *NameMap[V]) NewNameMapWithMap(nvmap orderedmap.OrderedMap[string, V]) {
	nm.nvmap = &nvmap
}

func (nm *NameMap[V]) NewNameMapEmpty() {
	nm.nvmap = orderedmap.New[string, V]()
}

// func (nm *NameMap[V]) HashCode() int {
// 	return nm.nvmap.hashcode()
// }
