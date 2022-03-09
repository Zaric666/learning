package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{11, 31, 22, 14, 4, 61, 13, 78, 21}
	b := quickSort(a, 0, len(a)-1)
	fmt.Println(b)
}
