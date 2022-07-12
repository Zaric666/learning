package a

import "fmt"

var Va2 = f("a2")

func init() {
	fmt.Println("func a2 init.  文件命名：a1.go先于a2.go")
}
