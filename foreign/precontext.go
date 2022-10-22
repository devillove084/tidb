package foreign

type VariableEnum = string

const (
	UTC_TIMESTAMP     VariableEnum = "utcTimestamp"
	CURRENT_TIMESTAMP VariableEnum = "currentTimestamp"
	LOCAL_TIMESTAMP   VariableEnum = "localTimestamp"
	CANCEL_FLAG       VariableEnum = "cancelFlag"
	TIMEOUT           VariableEnum = "timeout"
	SQL_ADVISOR       VariableEnum = "sqlAdvisor"
	STDERR            VariableEnum = "stderr"
	STDIN             VariableEnum = "stdin"
	STDOUT            VariableEnum = "stdout"
	LOCALE            VariableEnum = "locale"
	TIME_ZONE         VariableEnum = "timeZone"
	USER              VariableEnum = "user"
	SYSTEM_USER       VariableEnum = "systemUser"
)

type Variable struct {
	VariableEnum
	camelName string
	clazz     interface{}
}

func (v *Variable) Get(datacontext DataContext) interface{} {
	return nil
}

type DataContext interface {
	GetRootSchema() (SchemaPlus, error)
	GetTypeFactory() (GoTypeFactory, error)
	GetQueryProvider() (QueryProvider[interface{}], error)
	Get(name string) (interface{}, error)
}

type PreContext interface {
	GetTypeFactory() (GoTypeFactory, error)
	GetRootSchema() (ForeignSchema, error)
	GetMutableRootSchema() (ForeignSchema, error)
	GetDefaultSchemaPath() ([]string, error)
	Config() (ForeignConnectionConfig, error)
	GetDataContext() (DataContext, error)
	GetObjectPath() ([]string, error)
	GetRelRunner() (RelRunner, error)
}
