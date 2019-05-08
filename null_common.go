package typ

// Common general interface of accessor
type Common interface {
	Present() bool
	Valid() bool
	Typ(options ...Option) *Type
	Err() error
}
