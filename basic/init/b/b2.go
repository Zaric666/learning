package b

import (
	"fmt"
)

var Vb2 = f("b2")

func init() {
	fmt.Println("func b2 init. a包引入b包")
}
