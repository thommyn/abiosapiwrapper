package jsonconv

import "errors"

type rostersSubNode struct {
	subNodeName string
}

func NewRostersSubNode(subNodeName string) JsonConverter {
	return &rostersSubNode{
		subNodeName: subNodeName,
	}
}

func (c rostersSubNode) Convert(injson []interface{}) (outjson []interface{}, err error) {
	defer func() {
		// return error message if exception occurred reading roster nodes
		if recover() != nil {
			outjson = nil
			err = errors.New("unable to convert supplied json")
		}
	}()

	// TODO: Simplify code.......

	// array of {subNodeName} subnodes
	var arr []interface{}
	for _, serie := range injson {
		seriemap := serie.(map[string]interface{})
		if rostersNode, ok := seriemap["rosters"]; ok {
			rosters := rostersNode.([]interface{})
			for _, roster := range rosters {
				rostermap := roster.(map[string]interface{})
				if subnode, ok := rostermap[c.subNodeName]; ok {
					subnodes := subnode.([]interface{})
					for _, node := range subnodes {
						arr = append(arr, node)
					}
				}
			}
		}
	}

	return arr, nil
}
