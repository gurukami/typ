package typ

// NullCommon
type NullCommon interface {
	Present() bool
	Valid() bool
	Typ(options ...Option) *Type
	Err() error
}
