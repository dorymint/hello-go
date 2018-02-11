package reflect_test

import (
	"reflect"
	"testing"
)

func TestTypeSwitch(t *testing.T) {
	var integer int
	var i interface{}
	i = integer
	rt := reflect.TypeOf(i)
	switch i.(type) {
	case int:
		t.Log("i is int!")
	default:
		t.Log("i is unknown!")
	}
	t.Log("rt:", rt)
}
