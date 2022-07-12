package b

import (
	"fmt"
)

var Vb1 = f("b1")

func init() {
	fmt.Println("func b1 init. a包引入b包")
}

func f(b string) string {
	fmt.Println(b, "变量初始化")
	return b
}
