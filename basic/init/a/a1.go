package a

import (
	"fmt"
	_ "github.com/Zaric666/learning/basic/init/b"
)

var Va1 = f("a1")

func init() {
	fmt.Println("func a1 init. 文件命名：a0.go先于a1.go")
}
