package foreign

// type JsonSchemaInterface interface {
// 	Accept(handler ModelHandler)
// 	VisitChildren(moder ModelHandler)
// }

type JsonMaterialization struct {
	view           string
	table          string
	sql            interface{}
	viewSchemaPath []string
}

func (jm *JsonMaterialization) Accept(handler ModelHandler) {
	handler.VisitJsonMaterialization(jm)
}

type JsonMeasure struct {
	agg  string
	args interface{}
}

func (jme *JsonMeasure) Accept(handler ModelHandler) {
	handler.VisitJsonMeasure(jme)
}

type JsonTile struct {
	dimensions []interface{}
	measures   []JsonMeasure
}

func (jl *JsonTile) Accept(handler ModelHandler) {
	handler.VisitJsonTile(jl)
}

type JsonLattice struct {
	name               string
	sql                interface{}
	auto               bool
	algorithm          bool
	algorithmMaxMillis uint64
	statisticProvider  string
	tiles              []JsonTile
	defaultMeasures    []JsonMeasure
}

func (jl *JsonLattice) Accept(handler ModelHandler) {

}

type JsonSchema struct {
	name             string
	path             []interface{}
	materializations []JsonMaterialization
	lattices         []JsonLattice
	cache            bool
	autoLattice      bool
}

func (js *JsonSchema) NewJsonSchema(name string, path []interface{}, cache bool) *JsonSchema {
	return &JsonSchema{
		name:        name,
		path:        path,
		cache:       cache,
		autoLattice: false,
	}
}

func (js *JsonSchema) Accept(handler ModelHandler) {}

func (js *JsonSchema) VisitChildren(modelHandler ModelHandler) {
	for _, jl := range js.lattices {
		jl.Accept(modelHandler)
	}

	for _, jm := range js.materializations {
		jm.Accept(modelHandler)
	}
}

type JsonSchemaEnumtype int

const (
	MAP JsonSchemaEnumtype = iota
	GDBC
	CUSTOM
)

type JsonRoot struct {
	version       string
	defaultSchema string
	schemas       []JsonSchema
}

func (jr *JsonRoot) NewJsonRoot(version, defaultSchema string) JsonRoot {
	jr.schemas = nil
	jr.version = version
	jr.defaultSchema = defaultSchema
	return *jr
}
