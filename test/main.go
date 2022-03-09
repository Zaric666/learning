package main

import (
	"fmt"
	"os"
)

func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret := 10
		defer func() {
			ret = ret + 1
		}()
	}
	return
}

func main() {
	//fmt.Println(watShadowDefer(50))
	fmt.Println("func err2:", test2())
}

func test2() (err error) {

	defer func() {
		fmt.Println("defer err2:", err)
	}()

	if _, err := os.Open("xxx"); err != nil {
		return // return without err will compilation error
	}

	return
}
