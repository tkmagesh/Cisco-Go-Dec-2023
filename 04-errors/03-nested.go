package main

import (
	"errors"
	"fmt"
)

var ErrF1 = errors.New("f1 error")
var ErrF2 = errors.New("f2 error")

func main() {
	err := f1()
	if errors.Is(err, ErrF1) {
		fmt.Println("error occurred in f1")
	}
	if errors.Is(err, ErrF2) {
		fmt.Println("error occurred in f2")
	}
}

func f1() error {
	// return ErrF1
	e := f2()
	return fmt.Errorf("%w %w", ErrF1, e)
}

func f2() error {
	return ErrF2
}
