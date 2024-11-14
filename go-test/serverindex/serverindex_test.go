package serverindex

import (
	"testing"
)

// 基准测试
/* 这里用go test -v -bench=.运行所有基准测试，但是会出现no tests to run的问题
 * 在powershell中使用go test -v -bench.或者go test -v -bench="."
 * 在cmd中使用go test -v -bench.或者go test -v -bench="."  */
// https://www.cnblogs.com/linguoguo/p/10371253.html

// 这里基准测试函数主要测试了Select()函数在单线程和多线程的执行效率

func BenchmarkSelect(b *testing.B) {
	InitServerIndex()
	// 这里初始化计时器，确保只计算了Select()的执行时间
	b.ResetTimer()
	// b.N是基准测试框架自动调整的一个参数，用于控制基准测试函数的迭代次数
	for i := 0; i < b.N; i++ {
		// Select()
		FastSelect() // 优化版本
	}
}

// 并行基准测试函数
func BenchmarkSelectParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	// b.RunParallel会根据系统可用的处理器数量来分配工作（创建多个goroutine来执行参数的函数），
	// 每个goroutine都会执行Select()函数，直到所有的工作都完成。
	b.RunParallel(func(pb *testing.PB) {
		// 这是一个布尔值方法，返回true 表示还有更多的工作需要完成，
		// 返回 false 表示所有的工作都已经完成。
		// 这个循环会一直执行，直到 pb.Next() 返回 false。
		// 这意味着每个 goroutine 都会持续调用 Select()，直到所有的迭代次数都完成。
		for pb.Next() {
			// Select()
			FastSelect() //优化版本
		}
	})
}
