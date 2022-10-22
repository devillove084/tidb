package foreign

import "github.com/modern-go/reflect2"

type GoTypeFactory interface {
	RelDataTypeFactory
	CreateStructType(clazz interface{}) (RelDataType, error)
	CreateType(t reflect2.Type) (RelDataType, error)
	GetGoClass(t RelDataType) (reflect2.Type, error)

	CreateSyntheticType(types []reflect2.Type) (reflect2.Type, error)
	ToSql(t RelDataType) (RelDataType, error)
}
