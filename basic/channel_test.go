package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelBasic(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Second)
			ch <- i
		}
	}()

	for {
		select {
		case x := <-ch:
			fmt.Println("A收到", x)
		case y := <-ch:
			fmt.Println("B收到", y)
		default:

		}
	}
}

// 当 Channel 没有接收者能够处理数据时，向 Channel 发送数据会被下游阻塞
func TestBlockSend(t *testing.T) {
	ch := make(chan int)
	ch <- 1 // 向无缓冲区的channel发送数据，阻塞

	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3 // 向有缓冲区的channel发送数据，缓冲区满，阻塞
}

func TestBlockReceive(t *testing.T) {
	ch := make(chan int, 1)
	/*go func() {
		x := <-ch
		fmt.Println(x)
	}()*/
	ch <- 1

}
