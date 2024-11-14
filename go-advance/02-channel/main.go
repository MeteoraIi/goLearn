package main

import("fmt")

func CalSquare() {
	// 生产-消费， 同步问题，缓冲来缓和消费速度的问题
	src := make(chan int)
	dest := make(chan int, 3)
	// 这里用匿名函数是为了简洁性与可读性
	go func() {
		defer close(src)
		for i := 0;i < 10;i ++{
			// 无缓冲通道，会阻塞，直到有接受方从中接收数据
			src<-i
		}
	}()

	go func() {
		defer close(dest)
		for i := range src {
			// 带缓冲通道，会阻塞，直到缓冲区有空闲
			dest<- i*i
		}
	}()

	for i := range dest{
		fmt.Println(i)
	}
}

func main(){
	CalSquare()
}