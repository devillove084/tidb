package foreign

type ForeignConnection interface {
	GetRootSchema() (SchemaPlus, error)
	GetTypeFactory() (GoTypeFactory, error)
	//GetProperties() (Properties, error)
	SetSchema(schema string) error
	GetSchema() (SchemaPlus, error)

	Config()

	CreatePrepareContext() (PreContext, error)
}

type ForeignConnectionConfig interface {
	ForeignConnection
}

type ForeignStorage interface {
}

func Register(drivername string, driver Driver) {

}

type Driver interface {
	Open(path string) (ForeignStorage, error)
}

type ForeignDriver struct {
}

func (fd ForeignDriver) Open(path string) (ForeignStorage, error) {
	return nil, nil
}
