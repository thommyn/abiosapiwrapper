package jquery

const PathNodeSeparator = "/"

type JsonQueryFactory interface {
	Get(path string) JsonQuery
}

type defaultJsonQueryFactory struct {}

func NewDefaultJsonConverterFactory() JsonQueryFactory {
	return &defaultJsonQueryFactory{}
}

func (f defaultJsonQueryFactory) Get(path string) JsonQuery {
	return NewJsonQuery(path, PathNodeSeparator)
}
