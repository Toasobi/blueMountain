package main

import (
	"container/list"
	"fmt"
)

type MyQueue struct {
	l list.List
}

func Constructor() MyQueue {
	myQueue := &MyQueue{}
	return *myQueue
}

func (this *MyQueue) Push(x int) {
	this.l.PushBack(x)
}

func (this *MyQueue) Pop() interface{} {
	fmt.Println(this.l.Front().Value)
	value := this.l.Front().Value
	this.l.Remove(this.l.Front())
	return value
}

func (this *MyQueue) Peek() interface{} {
	return this.l.Front().Value
}

func (this *MyQueue) Empty() bool {
	return this.l.Len() == 0
}

func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(5)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)

}
