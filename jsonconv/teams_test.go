package jsonconv

import (
	"testing"
	"encoding/json"
)

func Test_LiveTeams_Convert_SimpleJson_ReturnPlayerNodesArrays(t *testing.T) {
	testJson :=
		`[{"id":1,"rosters":[{"id":10,"teams":[{"id":100},{"id":101}]},{"id":11,"teams":[{"id":110},{"id":111}]}]},
		{"id":2,"rosters":[{"id":20,"teams":[{"id":200},{"id":201}]},{"id":21,"teams":[{"id":210},{"id":211}]}]}]`
	expectedResult := `[{"id":100},{"id":101},{"id":110},{"id":111},{"id":200},{"id":201},{"id":210},{"id":211}]`

	conv := NewTeamsFromSeries()
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

func Test_LiveTeams_Convert_StrangeJson_ReturnError(t *testing.T) {
	testJson :=
		`[{"id":1}, {"id":2}]`

	conv := NewTeamsFromSeries()
	bytes := []byte(testJson)
	var data []interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		t.Errorf("Unable to unmarshal test json bytes. %s", err.Error())
	}

	res, err := conv.Convert(data)
	if err != nil {
		t.Errorf("An unexpected error occured when converting json. %s", err)
	}
	if len(res) != 0 {
		t.Errorf("Non zero array returned.")
	}
}

func Test_Convert_EmptyArrays_NoErrorReturned(t *testing.T) {
	testJson :=
		`[{"id":1,"rosters":[{"id":10,"teams":[]},{"id":11,"teams":[{"id":110},{"id":111}]}]},{"id":2,"rosters":[]}]`

	conv := NewPlayersFromSeries()
	bytes := []byte(testJson)
	var data []interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		t.Errorf("Unable to unmarshal test json bytes. %s", err.Error())
	}

	_, err = conv.Convert(data)
	if err != nil {
		t.Errorf("Unexpected error occured when converting empty nodes. %s", err.Error())
	}
}
