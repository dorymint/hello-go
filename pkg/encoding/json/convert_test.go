package json_test

// TODO: impl

import (
	"encoding/json"
	"fmt"
	"testing"
)

func ParseJSON(result *interface{}, i interface{}, t *testing.T) (err error) {
	fmt.Println("i:", i)
	if err != nil {
		return err
	}
	switch it := i.(type) {
	case map[string]interface{}:
		t.Log(it)
		for x, y := range it {
			t.Log(x, y)
		}
	default:
		err = fmt.Errorf("invalid types")
		t.Log("invalid types")
	}
	return err
}

func TestDecode(t *testing.T) {
	b := []byte(`{
	"hello": {
		"nest": true
	}
}`)
	var buf interface{}
	if err := json.Unmarshal(b, &buf); err != nil {
		t.Fatal(err)
	}
	t.Logf("buf: %#v\n", buf)

	var result *interface{}
	str := ParseJSON(result, buf, t)
	t.Log("str:", str)
}
