package basic

import (
	"fmt"
	"testing"
)

func TestNewSlice(t *testing.T) {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%-20p \n", &arr[1])

	// 左闭右开 得到新的slice [2,3]
	// cap 第三个参数指定则 5-2=3
	// cap 默认：数组长度-起始 10-2=8
	// s0=[2 3],len=2,cap=3
	s0 := arr[2:4:5] // low <= high <= max
	fmt.Printf("s0=%v,len=%d,cap=%d \n", s0, len(s0), cap(s0))
	// s0=[2 3],len=2,cap=8
	s0 = arr[2:4]
	fmt.Printf("s0=%v,len=%d,cap=%d \n", s0, len(s0), cap(s0))

	// 两个切片指向同一个数组
	s1 := arr[1:4]
	s2 := arr[1:5]
	// s1=[1 2 3],len=3,cap=9 0xc000014198
	fmt.Printf("s1=%v,len=%d,cap=%d %-20p \n", s1, len(s1), cap(s1), &s1[0])
	// s2=[1 2 3 4],len=4,cap=9 0xc000014198
	fmt.Printf("s2=%v,len=%d,cap=%d %-20p \n", s2, len(s2), cap(s2), &s2[0])
	// 其中一个修改数据，将影响指向的数组，并影响其他切片的值
	s1[0] = 97
	s1 = append(s1, 100, 101, 102, 103, 104, 105)
	// arr=[0 97 2 3 100 101 102 103 104 105] 0xc0000141b0
	fmt.Printf("arr=%v %-20p \n", arr, &arr[4])
	// s1=[97 2 3 100 101 102 103 104 105],len=9,cap=9 0xc0000141b0
	fmt.Printf("s1=%v,len=%d,cap=%d %-20p \n", s1, len(s1), cap(s1), &s1[3])
	// s2=[97 2 3 100],len=4,cap=9 0xc0000141b0
	fmt.Printf("s2=%v,len=%d,cap=%d %-20p \n", s2, len(s2), cap(s2), &s2[3])

	// 扩容切片，超出数组的范围，切片将指向新生成的底层数组
	s1 = append(s1, 106)
	// s1=[97 2 3 100 101 102 103 104 105 106],len=10,cap=18 0xc0001200a8
	fmt.Printf("s1=%v,len=%d,cap=%d %-20p \n", s1, len(s1), cap(s1), &s1[3])

	// arr=[0 97 2 3 100 101 102 103 104 105] 0xc0000141b0
	fmt.Printf("arr=%v %-20p \n", arr, &arr[4])
	// s2=[97 2 3 100],len=4,cap=9 0xc0000141b0
	fmt.Printf("s2=%v,len=%d,cap=%d %-20p \n", s2, len(s2), cap(s2), &s2[3])
}

/**
在分配内存空间之前需要先确定新的切片容量，运行时根据切片的当前容量选择不同的策略进行扩容：

如果期望容量大于当前容量的两倍就会使用期望容量；
如果当前切片的长度小于 1024 就会将容量翻倍；
如果当前切片的长度大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；
上述代码片段仅会确定切片的大致容量，下面还需要根据切片中的元素大小对齐内存，当数组中元素所占的字节大小为 1、8 或者 2 的倍数时，运行时会对齐内存：
*/
func TestAppendSlice(t *testing.T) {
	var arr []int64
	// 扩容 arr 切片并传入期望的新容量 5，这时期望分配的内存大小为 40 字节
	// 不过因为切片中的元素大小等于 sys.PtrSize，所以运行时会调用 runtime.roundupsize 向上取整内存的大小到 48 字节，所以新切片的容量为 48 / 8 = 6
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Printf("arr=%v,len=%d,cap=%d \n", arr, len(arr), cap(arr))
}

func TestReversToArray(t *testing.T) {
	s := make([]byte, 2, 4)
	s0 := (*[0]byte)(s)     // s0 != nil
	s1 := (*[1]byte)(s[1:]) // &s1[0] == &s[1]
	s2 := (*[2]byte)(s)     // &s2[0] == &s[0]
	//s4 := (*[4]byte)(s)    // panics: len([4]byte) > len(s)
	fmt.Println(s0, s1, s2)
}
