package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 10)
		ch <- "超时"
	}()
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-ticker.C:
			fmt.Println("ticker")
		}
	}
}
