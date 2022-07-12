package main

import (
	"fmt"
	_ "github.com/Zaric666/learning/basic/init/a"
)

func init() {
	fmt.Println("func main init")
}

func main() {
	fmt.Println("func main main")
}
