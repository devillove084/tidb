package foreign

type Wrapper interface {
	Unwrap(aClass interface{}) (interface{}, error)
	UnwrapOrError(aClass interface{}) (interface{}, error)
	MaybeUnwrap(aClass interface{}) (*interface{}, error)
}
