package main

import (
	"fmt"
)

type Counter interface {
	With() Counter
}

type A struct {
	I int
}

func (a *A) With() Counter {
	return a
}
func main() {
	var a1 Counter
	{
		var a A
		a1 = a.With()
	}
	fmt.Printf("%v", a1)
}
