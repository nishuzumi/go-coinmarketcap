package coinmarketcap

import (
	"encoding/json"
	"testing"
)

func TestStatus(t *testing.T) {
	data := []byte(`{"timestamp": "2018-06-02T22:51:28.209Z","error_code": 0,"error_message": "","elapsed": 10,"credit_count": 1}`)
	status := &Status{}
	err := json.Unmarshal(data, status)
	if err != nil {
		t.Fatal(err)
	}
}
