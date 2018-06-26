package jquery

const PathNodeSeparator = "/"

type JsonQueryFactory interface {
	Get(path string) JsonQuery
}

type jqueryJsonQueryFactory struct {}

func NewJqueryJsonConverterFactory() JsonQueryFactory {
	return &jqueryJsonQueryFactory{}
}

func (f jqueryJsonQueryFactory) Get(path string) JsonQuery {
	return NewJsonQuery(path, PathNodeSeparator)
}
