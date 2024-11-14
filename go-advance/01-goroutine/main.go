package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello goroutine :" + fmt.Sprint(i))
}

// "快速"打印
func main(){
	for i := 0;i < 5;i ++ {
		// goroutine是异步执行的
		// 如果使用hello(i)，可能会导致i直接迭代变成下一轮循环的值传入
		// 这里使用匿名函数来解决
		/* 闭包机制：
		 * 	   在 Go 语言中，匿名函数是一种闭包。闭包是一个函数值，它引用了其外部作用域中的变量。
		 *     当闭包被创建时，它会捕获外部作用域中的变量，形成一个“闭合”的环境。
		 * 变量捕获机制：
		 *     每个匿名函数实例都会捕获一个独立的 j 变量，这个 j 变量的值是在调用匿名函数时传递的 i 的值。
		 *     这里就是把i与j的值一一对应防止捕获出错
		 */
		go func(j int){
			hello(j)
		}(i)
	}

	// 同步
	time.Sleep(time.Second)
}