package main

import (
	"fmt"
	"time"
)

// chan 必须初始化，无论是做close信号的chan struct{}还是做消息传递的chan xxx，都需要make。这需要例子来明确。

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}

func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}

func main() {
	ch := make(chan int, 10)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}
