package foreign

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/edwingeng/deque/v2"
	"github.com/pingcap/tidb/util/logutil"
)

type Builder struct {
}

type TileBuilder struct{}

type ObjectMapper struct{}

func (om *ObjectMapper) ReadValue(uri string, root *JsonRoot) *JsonRoot {
	content, err := ioutil.ReadFile(uri)
	if err != nil {
		logutil.BgLogger().Error(err.Error())
	}
	err = json.Unmarshal(content, &root)
	if err != nil {
		logutil.BgLogger().Error(err.Error())
	}

	return root
}

type ModelHandler struct {
	connection     ForeignConnection
	latticeBuilder Builder
	tileBuilder    TileBuilder
	mapper         ObjectMapper
	modelUri       string
	schemaStack    deque.Deque[Pair[string, SchemaPlus]]
}

func (mh *ModelHandler) NewModelHandler(c ForeignConnection, uri string) {
	mh.connection = c
	mh.modelUri = uri

	res, err := mh.checkMapper(uri)
	if !res {
		panic(err)
	}
	mh.InitMappper()
	root := new(JsonRoot)
	v := mh.mapper.ReadValue(uri, root)
	mh.VisitJsonRoot(v)
}

func (mh *ModelHandler) checkMapper(uri string) (bool, error) {
	if !strings.Contains(uri, "json") {
		return false, nil
	}
	if _, err := os.Stat(uri); errors.Is(err, os.ErrNotExist) {
		return false, os.ErrNotExist
	}
	return true, nil
}

func (mh *ModelHandler) InitMappper() {

}

func (mh *ModelHandler) VisitJsonMaterialization(v *JsonMaterialization) {

}

func (mh *ModelHandler) VisitJsonRoot(root *JsonRoot) {
	schema_p, err := mh.connection.GetSchema()
	if err != nil {
		logutil.BgLogger().Panic(err.Error())
	}
	mh.schemaStack.PushBack(Pair[string, SchemaPlus]{"", schema_p})
	for _, schema := range root.schemas {
		schema.Accept(*mh)
	}

	mh.schemaStack.PopFront()
	if root.defaultSchema != "" {
		mh.connection.SetSchema(root.defaultSchema)
	}

}

func (mh *ModelHandler) VisitJsonSchema(root *JsonSchema) {
	// todo:
}

func (mh *ModelHandler) VisitJsonCustomSchema(root *JsonSchema) {
	// todo:
}

// func (mh *ModelHandler) VisitJsonGdbcSchema(root *JsonGdbcSchema) {

// }

func (mh *ModelHandler) VisitJsonTile(root *JsonTile) {

}

func (mh *ModelHandler) VisitJsonMeasure(root *JsonMeasure) {

}
