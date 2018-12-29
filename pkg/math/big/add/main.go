package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	i := big.NewInt(1)
	for {
		fmt.Println(i)
		time.Sleep(time.Second)
		i = i.Add(i, i)
	}
}
