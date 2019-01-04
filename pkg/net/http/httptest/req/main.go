package main

import (
	"fmt"
	"net/http/httptest"
)

func main() {
	req := httptest.NewRequest("GET", "http://localhost:8080/api/random", nil)
	fmt.Println(req.URL.Path)
}
