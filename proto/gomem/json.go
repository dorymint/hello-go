package gomem

import (
	"encoding/json"
	"fmt"
	"log"
)

type gomemJSON struct {
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Date    string   `json:"date"`
	Tag     []string `json:"tag"`
}

// JSON check
func JSON() {
	var gj []gomemJSON
	b := []byte(`[
	{
		"name":    "todo",
		"content": "hi",
		"date":    "20013030",
		"tag":     ["todo","next"]
	},
	{
		"name":    "next",
		"content": "hello",
		"date":    "20120303",
		"tag":     ["todo","next"]
	}
	]`)
	if err := json.Unmarshal(b, &gj); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", gj)
	fmt.Println("prot")
}
