package jsonconv

import (
	"strings"
	"errors"
)

type JsonQuery interface {
	GetSubNodes(jsonData []interface{}) (nodes []interface{}, err error)
}

type jsonQuery struct {
	path string
	separator string
}

func NewJsonQuery(path string, separator string) *jsonQuery {
	return &jsonQuery{
		path: path,
		separator: separator,
	}
}

func (q jsonQuery) GetSubNodes(jsonData []interface{}) (nodes []interface{}, err error) {
	return q.getSubNodes(q.path, jsonData)
}

func (q jsonQuery) getSubNodes(path string, data []interface{}) (nodes []interface{}, err error) {
	paths := strings.SplitN(path, q.separator, 2)

	// if last node in path...
	if len(paths) == 1 {
		if paths[0] != "*" {
			return nil, errors.New("last path must be *")
		}
		return data, nil
	}

	// loop nodes recursive
	for _, d := range data {
		dmap := d.(map[string]interface{})
		if dmapnode, ok := dmap[paths[0]]; ok {
			subnodes := dmapnode.([]interface{})
			subnodesarr, err := q.getSubNodes(paths[1], subnodes)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, subnodesarr...)
		}
	}

	return nodes, nil
}
