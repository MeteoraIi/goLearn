package main

import (
	"fmt"
	"sync"
	"time"
)

// 声明全局变量
var(
	x int64
	lock sync.Mutex
)

func addWithLock() {
	for i := 0;i < 2000;i ++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func addWithoutLock() {
	for i := 0;i < 2000;i ++ {
		// 此处没有加锁，会出现5个协程同时进入，最后x被赋值是都有可能的
		x += 1
	}
}

func main(){
	x = 0
	// 五个协程并发执行
	for i := 0;i < 5;i ++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	
	fmt.Println("WithoutLock", x)

	x = 0
	for i := 0;i < 5;i ++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	
	fmt.Println("WithoutLock", x)
}
