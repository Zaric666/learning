package basic

import (
	"fmt"
	"testing"
)

func DeferFunc1(i int) (t int) { // 1.入参1，返回值赋值t=0
	t = i          // 2.t=1
	defer func() { // 3.压栈
		t += 3 // 5.修改t = 1+3 = 4 方法执行完毕，t=4
	}()
	return t // 4.返回t
}

func DeferFunc2(i int) int { // 1.入参1，注意这个返回匿名变量
	t := i         // 2.t赋值i = 1
	defer func() { // 3.入栈
		t += 3 // 5.赋值 t = 1+3 = 4
	}()
	return t // 4.返回1(并不是返回t) 可以简单理解为将t的值拷贝给了匿名的返回变量
}

/**
也可以按照如下代码理解
func DeferFunc2(i int) (result int) {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
上面的代码return的时候相当于将t赋值给了result，当defer修改了t的值之后，对result是不会造成影响的
**/

func DeferFunc3(i int) (t int) { // 1.入参i=1,返回值t=0
	defer func() { // 2.压栈
		t += i // 4.修改t = 2+1 = 3 最终方法执行完毕 t=3
	}()
	return 2 // 3.赋值t=2 返回t
}

func DeferFunc4() (t int) { // 1.初始化返回值t为零值 0
	defer func(i int) { // 2.defer 压栈 入参 t为0
		fmt.Println(i) // 5.打印入参 （第2步 入参t=0) t->i 0
		fmt.Println(t) // 6.打印 t(第4步 t赋值2) 2
	}(t)
	t = 1    // 3.t赋值1
	return 2 // 4.t赋值2
}

func TestDefer(t *testing.T) {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	DeferFunc4()
}
