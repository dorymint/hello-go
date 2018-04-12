package fmt_test

import (
	"fmt"
	"testing"
)

func TestTab(t *testing.T) {
	t.Log(fmt.Sprintf("hello\thello world"))
	const indent = "\t"
	s := fmt.Sprintf("hello%shello world", indent)
	t.Log(s)
}
