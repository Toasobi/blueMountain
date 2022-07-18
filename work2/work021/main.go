package main

import (
	"fmt"
	"sync"
)

//简单生产者消费者模型
var wg sync.WaitGroup

//生产者
func create(x int, ch chan<- int) {
	for i := 0; i < x; i++ {
		fmt.Printf("生产者生产了第%d件商品\n", i)
		ch <- i
	}
	close(ch)
	wg.Done()
}

//消费者
func custom(ch <-chan int) {
	for v := range ch {
		fmt.Printf("消费者消费了第%d件商品\n", v)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	ch1 := make(chan int)
	var a int = 10
	go create(a, ch1)
	go custom(ch1)
	wg.Wait()
}
