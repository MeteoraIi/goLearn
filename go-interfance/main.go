package main

import "fmt"

// 实现了这个方法，就说明我属于这个接口
type Animal interface {
	Eat()
	Run()
}

type Cat struct {
	Name string
	Sex  bool
}

type Dog struct {
	Name string
}

// o(=•ェ•=)实现了这个接口
// 方法: 绑定到这个参数对应的类型
// c == this， go中由自己来指定接收者的名字
// 所以Cat是实现了接口的类型
// go可以将方法绑定到任何类型上
func (c Cat) Eat() {
	fmt.Println(c.Name, "开始吃")
}
func (c Cat) Run() {
	fmt.Println(c.Name, "开始跑")
}

// U•ェ•*U实现了吗
func (c Dog) Eat() {
	fmt.Println(c.Name, "开始吃")
}
func (c Dog) Run() {
	fmt.Println(c.Name, "开始跑")
}

// 泛型，想传什么传什么
// 泛型允许程序员在定义函数、结构体、接口等时使用类型参数，而不是具体的类型。
// func MyFun(a interfance{}) {}

//
/* func MyFun(a Animal) {
 *	a.Run()
 *	a.Eat()
 *}
 */

// 可以更灵活的在任何地方使用L来调用Animal接口的方法
// 可以在逻辑不同的情况下调用L的方法
// 自由度大

// L不管在什么地方调用都能run起来
// 这里的L主要是在代码逻辑层面提供了一种相对集中的调用入口和定位修改点的方式。
var L Animal

func MyFun(a Animal) {
	L = a
}

func main() {
	//c := Dog{
	//	"Spake",
	//}
	//
	//d := Cat{
	//	"Tom",
	//	false,
	//}
	//
	//MyFun(c)
	//MyFun(d)

	c := Cat{
		"Tom",
		false,
	}

	MyFun(c)
	L.Run()
	L.Eat()
}
