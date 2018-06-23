package jsonconv

import (
	"fmt"
)

type JsonConverter interface {
	Convert(injson []interface{}) (outjson []interface{}, err error)
}

type JsonConverterFactory interface {
	Get(convType string) (JsonConverter, error)
}

type defaultJsonConverterFactory struct {}

func NewDefaultJsonConverterFactory() JsonConverterFactory {
	return &defaultJsonConverterFactory{}
}

func (f defaultJsonConverterFactory) Get(convType string) (JsonConverter, error) {
	switch convType {
	case "none":
		return nil, nil
	case "liveplayers":
		return NewPlayersFromSeries(), nil
	case "liveteams":
		return NewTeamsFromSeries(), nil
	default:
		return nil, fmt.Errorf("unknown converter type %s", convType)
	}
}
