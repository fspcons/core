package mocks

type SqlFilterable struct{}

func (ref SqlFilterable) GetFilter() map[string]any {
	return map[string]any{}
}

type NoSqlFilterable struct{}

func (ref NoSqlFilterable) GetFilter() map[string]any {
	return map[string]any{}
}

type updateInputType interface {
	RefreshTimestamp() updateInputType
	Validate() error
}
