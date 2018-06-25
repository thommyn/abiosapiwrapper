package jsonconv

const PathNodeSeparator = "/"

type JsonQueryFactory interface {
	Get(convType string) JsonQuery
}

type defaultJsonQueryFactory struct {}

func NewDefaultJsonConverterFactory() JsonQueryFactory {
	return &defaultJsonQueryFactory{}
}

func (f defaultJsonQueryFactory) Get(path string) JsonQuery {
	return NewJsonQuery(path, PathNodeSeparator)
}
