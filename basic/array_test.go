package basic

import (
	"fmt"
	"testing"
)

func TestArrayBasic(t *testing.T) {
	arr := [3]int{1, 2, 3}
	arr1 := [3]int{1, 2, 3}
	// 数组可以用 == != 比较，仅限元素同类型的数组对比
	if arr == arr1 {
		fmt.Println("arr equal to arr1")
	}
	// mismatched types error
	/*arr2 := [3]int64{1,2,3}
	if arr != arr2 {
		fmt.Printf("arr not equal to arrr2")
	}*/
}
