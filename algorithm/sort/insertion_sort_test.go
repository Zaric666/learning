package sort

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{11, 31, 20, 14, 4, 61, 13, 78, 21}
	c := []int{4, 11, 13, 14, 20, 21, 31, 61, 78}
	b := insertSort(a)
	assert.Equal(t, c, b)
}
