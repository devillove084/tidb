package foreign

type DefaultEdge struct {
	source interface{}
	target interface{}
}

type IntPair struct {
	source int
	target int
}

type Step struct {
	DefaultEdge
	keys      []IntPair
	keyString string
}

type Path struct {
	steps []Step
	id    int
}
