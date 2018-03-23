package base

import (
	"encoding/json"
	"testing"
)

func TestBase(t *testing.T) {
	type JSON struct {
		First string `json:"first"`
		Object []string `json:"object"`
	}
	var v JSON
	b := []byte(`{
	"first": "string",
	"object": [
		"value1",
		"value2"
	]
}`)

	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", v)
}
