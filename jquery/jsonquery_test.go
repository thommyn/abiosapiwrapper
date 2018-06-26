package jquery

import (
	"testing"
	"encoding/json"
	"errors"
	"fmt"
)

func Test_GetSubNodes_PlayersJson_ReturnPlayerNodesArray(t *testing.T) {
	testJson :=
		`[{"id":1,"rosters":[{"id":10,"players":[{"id":100},{"id":101}]},{"id":11,"players":[{"id":110},{"id":111}]}]},
		{"id":2,"rosters":[{"id":20,"players":[{"id":200},{"id":201}]},{"id":21,"players":[{"id":210},{"id":211}]}]}]`
	expectedResult := `[{"id":100},{"id":101},{"id":110},{"id":111},{"id":200},{"id":201},{"id":210},{"id":211}]`

	jquery := NewJsonQuery("rosters/players/*", "/")
	subNodesJson, err := getSubNnodesFromJson(testJson, jquery)
	if err != nil {
		t.Errorf("An unexpected error occured when querying json. %s", err)
	}

	if  subNodesJson != expectedResult {
		t.Errorf(
			"Json does not match the expected result, got: \"%s\", want: \"%s\".",
			subNodesJson,
			expectedResult)
	}
}

func Test_GetSubNodes_InvalidJson_ReturnNull(t *testing.T) {
	testJson :=	`[{"id":1}, {"id":2}]`
	var expectedResult string = "null"

	jquery := NewJsonQuery("rosters/players/*", "/")
	subNodesJson, err := getSubNnodesFromJson(testJson, jquery)
	if err != nil {
		t.Errorf("An unexpected error occured when querying json. %s", err)
	}

	if  subNodesJson != expectedResult {
		t.Errorf(
			"Json does not match the expected result, got: \"%s\", want: \"%s\".",
			subNodesJson,
			expectedResult)
	}
}

func Test_GetSubNodes_EmptyArrays_ReturnCorrectNodes(t *testing.T) {
	testJson :=
		`[{"id":1,"rosters":[{"id":10,"players":[]},{"id":11,"players":[{"id":110},{"id":111}]}]},{"id":2,"rosters":[]}]`
	expectedResult := `[{"id":110},{"id":111}]`

	jquery := NewJsonQuery("rosters/players/*", "/")
	subNodesJson, err := getSubNnodesFromJson(testJson, jquery)
	if err != nil {
		t.Errorf("An unexpected error occured when querying json. %s", err)
	}

	if  subNodesJson != expectedResult {
		t.Errorf(
			"Json does not match the expected result, got: \"%s\", want: \"%s\".",
			subNodesJson,
			expectedResult)
	}
}

func getSubNnodesFromJson(jsondata string, jquery JsonQuery) (outjson string, err error) {
	bytes := []byte(jsondata)
	var data []interface{}

	if e := json.Unmarshal(bytes, &data); e != nil {
		err = errors.New(fmt.Sprintf("Unable to unmarshal test json bytes. %s", e.Error()))
		return
	}

	subNodes, e := jquery.GetSubNodes(data)
	if e != nil {
		err = errors.New(fmt.Sprintf("Unable to convert test json data. %s", e.Error()))
		return
	}

	subNodesBytes, e := json.Marshal(subNodes)
	if e != nil {
		err = errors.New(fmt.Sprintf("Unable to marshal test json data. %s", e.Error()))
		return
	}

	outjson = string(subNodesBytes)
	return
}
