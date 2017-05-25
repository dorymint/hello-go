package benchmap

import (
	"fmt"
	"os"
)

func BenchMapBool(s string, m map[string]bool) error {
	if m[s] {
		return fmt.Errorf("invalid")
	}
	return nil
}

func BenchMapInfo(s string, m map[string][]os.FileInfo) error {
	if _, ok := m[s]; ok {
		return fmt.Errorf("invalid")
	}
	return nil
}
