package jquery

import (
	"testing"
	"encoding/json"
)

func Test_Convert_SimpleJson_ReturnPlayerNodesArrays(t *testing.T) {
	testJson :=
		`[{"id":1,"rosters":[{"id":10,"players":[{"id":100},{"id":101}]},{"id":11,"players":[{"id":110},{"id":111}]}]},
		{"id":2,"rosters":[{"id":20,"players":[{"id":200},{"id":201}]},{"id":21,"players":[{"id":210},{"id":211}]}]}]`
	expectedResult := `[{"id":100},{"id":101},{"id":110},{"id":111},{"id":200},{"id":201},{"id":210},{"id":211}]`

	jquery := NewJsonQuery("rosters/players/*", "/")
	bytes := []byte(testJson)
	var data []interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		t.Errorf("Unable to unmarshal test json bytes. %s", err.Error())
	}

	subNodes, err := jquery.GetSubNodes(data)
	if err != nil {
		t.Errorf("Unable to convert test json data. %s", err.Error())
	}

	subNodesBytes, err := json.Marshal(subNodes)
	if err != nil {
		t.Errorf("Unable to marshal test json data. %s", err.Error())
	}

	if  subNodesString := string(subNodesBytes); subNodesString != expectedResult {
		t.Errorf(
			"Converted json does not match the expected result, got: \"%s\", want: \"%s\".",
			subNodesString,
			expectedResult)
	}
}

func Test_Convert_StrangeJson_ReturnEmptyArray(t *testing.T) {
	testJson :=
		`[{"id":1}, {"id":2}]`

	conv := NewJsonQuery("rosters/players/*", "/")
	bytes := []byte(testJson)
	var data []interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		t.Errorf("Unable to unmarshal test json bytes. %s", err.Error())
	}

	res, err := conv.GetSubNodes(data)
	if err != nil {
		t.Errorf("An unexpected error occured when converting json. %s", err)
	}
	if len(res) != 0 {
		t.Errorf("Non zero array returned.")
	}
}

func Test_Players_Convert_EmptyArrays_NoErrorReturned(t *testing.T) {
	testJson :=
		`[{"id":1,"rosters":[{"id":10,"players":[]},{"id":11,"players":[{"id":110},{"id":111}]}]},{"id":2,"rosters":[]}]`

	conv := NewJsonQuery("rosters/players/*", "/")
	bytes := []byte(testJson)
	var data []interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		t.Errorf("Unable to unmarshal test json bytes. %s", err.Error())
	}

	_, err = conv.GetSubNodes(data)
	if err != nil {
		t.Errorf("Unexpected error occured when converting empty nodes. %s", err.Error())
	}
}
