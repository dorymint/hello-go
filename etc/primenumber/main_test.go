package main

import (
	"math/big"
	"reflect"
	"testing"
)

var trailDivTests = []struct {
	count string
	exp   []string
}{
	{"10", []string{"2", "3", "5", "7"}},
	{"20", []string{"2", "3", "5", "7", "11", "13", "17", "19"}},
}

func TestTrailDiv(t *testing.T) {
	for _, test := range trailDivTests {
		i, ok := new(big.Int).SetString(test.count, 10)
		if !ok {
			t.Fatalf("case %v can not set to big.Int", test)
		}
		rec, err := PrimeNumbers(i, "trail")
		if err != nil {
			t.Fatalf("case %+v err %v", test, err)
		}
		var out []string
		for bi := range rec {
			out = append(out, bi.String())
		}
		if !reflect.DeepEqual(test.exp, out) {
			t.Fatalf("exp: %+v but out: %+v", test.exp, out)
		}
	}
}
