package main

import (
	"fmt"
	"testing"
)

// go tool compile -m -l escape.go
// 内存分配逃逸到堆中，增加垃圾回收(GC)的负担。
// 在对象频繁创建和删除的场景下，传递指针导致的GC 开销可能会严重影响性能

// 给一个引用类对象中的引用类成员进行赋值，可能出现逃逸现象
// Go语言中的引用类型有func（函数类型），interface（接口类型），
// slice（切片类型），map（字典类型），channel（管道类型），*（指针类型）等
func TestReference(t *testing.T) {
	// []interface{}数据类型，通过[]赋值必定会出现逃逸
	data := []interface{}{100, 200}
	data[0] = 100

	// map[string]interface{}类型尝试通过赋值，必定会出现逃逸
	data1 := make(map[string]interface{})
	data1["key"] = 200

	// map[interface{}]interface{}类型尝试通过赋值，会导致key和value的赋值，出现逃逸
	data2 := make(map[interface{}]interface{})
	data2[100] = 200

	// map[string][]string数据类型，赋值会发生[]string发生逃逸
	data3 := make(map[string][]string)
	data3["key"] = []string{"value"}

	// []*int数据类型，赋值的右值会发生逃逸现象
	a := 10
	data4 := []*int{nil}
	data4[0] = &a
}

// channel
func TestChannel(t *testing.T) {
	// chan []string数据类型，想当前channel中传输[]string{"value"}会发生逃逸现象
	ch := make(chan []string)
	s1 := []string{"test"}
	go func() {
		ch <- s1
	}()

	// 向 channel 发送指针数据。因为在编译时，不知道channel中的数据会被哪个 goroutine
	// 接收，因此编译器没法知道变量什么时候才会被释放，因此只能放入堆中。
	ch1 := make(chan *int, 1)
	y := 5
	ch1 <- &y // y逃逸，因为y地址传入了chan中，编译时无法确定什么时候会被接收，所以也无法在函数返回后回收y
}

// 切片扩容后长度太大，导致栈空间不足，逃逸到堆上
func TestSlice(t *testing.T) {
	s := make([]int, 10000, 10000)
	for index, _ := range s {
		s[index] = index
	}
}

type Animal interface {
	speak()
}

type Dog struct {
}

func (dog Dog) speak() {

}

// 在 interface 类型上调用方法。 在 interface 类型上调用方法时
// 会把interface变量使用堆分配， 因为方法的真正实现只能在运行时知道
func TestInterface(t *testing.T) {
	var animal Animal
	animal = Dog{}
	animal.speak() // 调用方法时，Dog{} 发生逃逸，因为方法是动态分配的
}

// 函数中出现内存逃逸
func TestFunc(t *testing.T) {
	b := foo()
	fmt.Println(b)

	b1 := 1
	foo1(&b1)

	b2 := []string{"hello"}
	foo2(b2)
}

// 【指针逃逸】
// 局部变量做返回值被函数外部引用
func foo() *int {
	a := 1
	return &a
}

// func(*int)函数类型，进行函数赋值，会使传递的形参出现逃逸现象
func foo1(a *int) {
	return
}

// func([]string): 函数类型，进行[]string{"value"}赋值，会使传递的参数出现逃逸现象
func foo2(a []string) {
	return
}
