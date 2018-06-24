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

	var arr []interface{}
	for _, serie := range injson {

		// loop rosters array
		rosters := serie.(map[string]interface{})["rosters"].([]interface{})
		for _, roster := range rosters {

			// loop players
			subNodeArray := roster.	(map[string]interface{})[c.subNodeName].([]interface{})
			for _, subNode := range subNodeArray {

				// append player to array
				arr = append(arr, subNode)
			}
		}
	}

	return arr, nil
}
