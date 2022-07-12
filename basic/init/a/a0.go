package a

import (
	"fmt"
)

var a0 = f("a0")

func init() {
	fmt.Println("func a0 init. main包引入a包")
}

func f(a string) string {
	fmt.Println(a, "变量先初始化")
	return a
}
