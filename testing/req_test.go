package testing

import (
	"encoding/json"
	"strings"
	"testing"
)


func TestDecodeData(t *testing.T) {
	jsonData := `{"sum": 100}`
	jsonDataReader := strings.NewReader(jsonData)
	decoder := json.NewDecoder(jsonDataReader)
	var profile map[string]interface{}
	err := decoder.Decode(&profile)
	if err != nil {
		t.Errorf("We have error for Decode data")
	}
}