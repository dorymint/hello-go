package gomem

import (
	"testing"
)

func TestAddNum(t *testing.T) {
	x, y := 1, 2
	expected := 3
	if out := AddNum(x, y); out != expected {
		t.Fatalf("Expected %v, but %v:", expected, out)
	}
}

// Add
// List

