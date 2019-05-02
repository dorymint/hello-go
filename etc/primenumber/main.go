// etc/primenumber.
// count prime numbers.
//
//	go run main.go ALGORITHM TARGETNUMBER
//
package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"
)

type Checker interface {
	IsPrime(*big.Int) bool
}

type TrailDiv struct {
	// cache
	rem *big.Int

	// constant
	two *big.Int
}

func NewTrailDiv() *TrailDiv { return &TrailDiv{rem: new(big.Int), two: big.NewInt(2)} }

func (t *TrailDiv) IsPrime(i *big.Int) bool {
	if i.Sign() != 1 {
		return false
	}
	if i.String() == "2" {
		return true
	}
	for j := big.NewInt(3); j.Cmp(i) == -1; j.Add(j, t.two) {
		if t.rem.Rem(i, j); t.rem.String() == "0" {
			return false
		}
	}
	return true
}

func PrimeNumbers(tar *big.Int, algo string) (<-chan *big.Int, error) {
	var checker Checker
	switch algo {
	case "trail":
		checker = NewTrailDiv()
	default:
		return nil, fmt.Errorf("unexpected algorithm: %s", algo)
	}

	two := big.NewInt(2)
	rec := make(chan *big.Int)

	// infinite
	if tar == nil {
		go func() {
			rec <- big.NewInt(2)
			for i := big.NewInt(3); true; i.Add(i, two) {
				if !checker.IsPrime(i) {
					continue
				}
				rec <- new(big.Int).Set(i)
			}
		}()
		return rec, nil
	}

	if tar.Sign() != 1 || tar.String() == "1" {
		return nil, fmt.Errorf("invalid target number: %s", tar.String())
	}
	go func() {
		rec <- big.NewInt(2)
		for i := big.NewInt(3); i.Cmp(tar) == -1; i.Add(i, two) {
			if !checker.IsPrime(i) {
				continue
			}
			rec <- new(big.Int).Set(i)
		}
		close(rec)
	}()
	return rec, nil
}

func run(args []string) error {
	var i *big.Int
	var algo string
	switch len(args) {
	case 0:
		fallthrough
	case 1:
		return errors.New("not enough arguments")
	case 2:
		algo = args[0]
		if args[1] == "infinite" {
			break
		}
		var ok bool
		i, ok = new(big.Int).SetString(args[1], 10)
		if !ok {
			return fmt.Errorf("can not convert: %s", args[1])
		}
	default:
		return fmt.Errorf("invalid arguments: %v", args)
	}

	rec, err := PrimeNumbers(i, algo)
	if err != nil {
		return err
	}
	for s := range rec {
		_, err := fmt.Println(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
