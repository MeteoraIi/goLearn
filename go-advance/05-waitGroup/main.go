package main

import (
	"fmt"
	"sync"

)

func hello(i int) {
	fmt.Println("hello goroutine :" + fmt.Sprint(i))
}

func MangGoWait() {
	var wg sync.WaitGroup
	// 在goroutine启动之前，传入要等待的goroutine数量
	wg.Add(5)

	for i := 0;i < 5;i ++ {
		go func(j int) {
			// 完成一个goroutine后减少计数器
			defer wg.Done()
			hello(j)
		}(i)
	}

	// 阻塞主goroutine，直到WaitGroup的计数器为0
	wg.Wait()
}

// "快速"打印
func main(){
	MangGoWait()
}