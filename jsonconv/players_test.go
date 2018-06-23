package jsonconv

import (
	"testing"
	"encoding/json"
)

func Test_Convert_SimpleJson_ReturnPlayerNodesArrays(t *testing.T) {
	testJson :=
		`[{"id":1,"rosters":[{"id":10,"players":[{"id":100},{"id":101}]},{"id":11,"players":[{"id":110},{"id":111}]}]},
		{"id":2,"rosters":[{"id":20,"players":[{"id":200},{"id":201}]},{"id":21,"players":[{"id":210},{"id":211}]}]}]`
	expectedResult := `[{"id":100},{"id":101},{"id":110},{"id":111},{"id":200},{"id":201},{"id":210},{"id":211}]`

	conv := NewPlayersFromSeries()
	bytes := []byte(testJson)
	var data []interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		t.Errorf("Unable to unmarshal test json bytes. %s", err.Error())
	}

	convdata, err := conv.Convert(data)
	if err != nil {
		t.Errorf("Unable to convert test json data. %s", err.Error())
	}

	convdatabytes, err := json.Marshal(convdata)
	if err != nil {
		t.Errorf("Unable to marshal test json data. %s", err.Error())
	}

	if  convdatajson := string(convdatabytes); convdatajson != expectedResult {
		t.Errorf(
			"Converted json does not match the expected result, got: \"%s\", want: \"%s\".",
			convdatajson,
			expectedResult)
	}
}

func Test_Convert_StrangeJson_ReturnError(t *testing.T) {
	testJson :=
		`[{"id":1}, {"id":2}]`

	conv := NewPlayersFromSeries()
	bytes := []byte(testJson)
	var data []interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		t.Errorf("Unable to unmarshal test json bytes. %s", err.Error())
	}

	if _, err := conv.Convert(data); err == nil {
		t.Errorf("Converting an invalid json does not return an error message.")
	}
}
