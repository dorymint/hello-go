package s

import (
	"testing"
)

func TestSimple(t *testing.T) {
	switch{
	case false, true, false:
		t.Log("false, true, false")
	default:
		t.Log("default")
	}
}
