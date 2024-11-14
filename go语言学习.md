# go语言学习

## 一、基础语法

### <1.>变量与常量

​	**golang中的变量是需要声明的**。

​	关键字**var**，`var a = "initial"`自动推导类型。也可以显式指出，`var a string = "initial"`

​	关键字**const**， `const a = "initial"`自动推导类型。也可以显式指出，`var a string = "initial"`

### <2.>"="与":="的区别

`:=`**（短变量声明）**

- **变量声明和初始化同时进行**：`:=`用于声明新变量并同时对其进行初始化。例如，`a := 10`声明了一个名为`a`的变量，并将其初始化为`10`。编译器会根据赋值的值自动推断变量的类型，在这里`a`被推断为`int`类型。
- **只能在函数内部使用**：这是一个重要的限制。`:=`操作符不能用于函数外部的变量声明。因为函数外部的变量声明通常涉及到包级别的变量，需要使用`var`关键字来明确声明变量的类型、作用域等信息。
- **新变量的创建**：如果使用`:=`声明变量，并且在当前作用域中已经存在同名的变量，编译器会认为你想要创建一个新的变量，而不是对已存在的变量进行赋值。例如：

​	``func main() {`
​    	`a := 10`
​    	`fmt.Println("a:", a)`
​    	`{`
​        	`a := 20`
​       	 `fmt.Println("inner a:", a)`
   	 `}`
   	 `fmt.Println("a:", a)`
​	}`

​	在这个例子中，内部代码块中`a := 20`创建了一个新的变量`a`，它的作用域仅限于内部代码块。当内部代码块执行完毕后，外部的`a`变量的值仍然是`10`。

**`=`（赋值操作）**

- **变量必须先声明**：与`:=`不同，`=`用于给**已经声明的变量赋值**。例如，`var b int`先声明了一个`int`类型的变量`b`，然后可以使用`b = 20`来给`b`赋值。
- **可用于函数内外**：`=`操作符可以在函数内部和外部使用。在函数外部，它用于给包级别的变量赋值；在函数内部，它用于给已经声明的函数内变量赋值。例如，在包级别声明变量：

**短变量（用:=声明的变量）：**自动类型推断， 局部作用域声明。

### <3.>switch

**特点：**

1、一个`case`语句可以匹配多个值。例如：

```go
day := "Monday"
switch day {
case "Saturday", "Sunday":
    println("It's the weekend")
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    println("It's a weekday")
default:
    println("Invalid day")
}
```

2、无表达式的switch（类似if - else if - else链）：

```go
score := 85
switch {
case score >= 90:
    println("A grade")
case score >= 80:
    println("B grade")
case score >= 70:
    println("C grade")
case score >= 60:
    println("D grade")
default:
    println("F grade")
}
```

3、类型断言（在接口类型转换中很有用）

`switch`还可以用于类型断言。例如，当有一个接口类型的值，想要确定它的具体实现类型时，可以使用`switch`。假设`i`是一个接口类型的值，并且可能是`int`或者`string`类型的实现：

```go
func printValue(i interface{}) {
    switch v := i.(type) {
    case int:
        println("It's an integer:", v)
    case string:
        println("It's a string:", v)
    default:
        println("Unknown type")
    }
}
```

### <4.>切片

1. **切片类型声明包含元素类型信息**

   ​	在 Go 语言中，当声明一个切片时，`[]`语法中的元素类型就确定了这个切片是哪种类型的切片。例如，`[]string`就是声明了一个字符串类型的切片，`[]int`是整数类型的切片。这意味着你已经明确了切片中元素所属的类型，而**不需要像在某些其他语言中额外说明** “这个切片是属于某个特定结构体或者对象的”。

   ​	切片是一种**引用类型**，它本身**不存储具体的数据元素**（它存储的是指向底层数组的指针、长度和容量信息），它只是对底层数组的一个 “视图”。所以重点在于切片中元素的类型，而不是所谓 “属于谁的切片”。

2. **切片的创建**

   - **使用`make`函数创建：**

     ​	语法为`make([]T, length, capacity)`，其中`T`是切片元素的类型，`length`是切片的初始长度，`capacity`是切片的初始容量。例如，`s := make([]int, 3, 5)`创建了一个`int`类型的切片`s`，其初始长度为 3，初始容量为 5。这意味着底层数组有 5 个元素的空间，但切片目前只包含 3 个元素。如果省略容量参数，如`make([]int, 3)`，则容量和长度相等，都是 3。

   - **从数组创建切片：**

     ​	**可以通过对数组进行切片操作来创建切片**。例如，`arr := [5]int{1, 2, 3, 4, 5}`，`s := arr[1:3]`创建了一个切片`s`，它引用了`arr`数组中索引为 1 和 2 的元素（在 Go 语言中，切片的区间是左闭右开的，即包含起始索引对应的元素，不包含结束索引对应的元素）。此时，`s`的长度为 2（因为包含两个元素），容量为 4（因为从索引 1 开始，到数组末尾还有 4 个元素）。

   - **切片的类型是[]type，而数组的类型是type**

3. **切片的操作**

- 切片的追加元素：

  ​        使用`append`函数来追加元素。例如，`s := make([]int, 3)`，`s = append(s, 4)`会将元素 4 追加到切片`s`的末尾。**如果切片的容量不足以容纳新元素，`append`函数会自动分配一个新的底层数组，并将原切片的元素和新元素复制到新数组中**。例如，若初始切片`s`的容量为 3，当追加一个元素时，Go 可能会创建一个新的底层数组，容量可能是原来的两倍（具体增长策略由 Go 语言实现决定），并将原切片元素和新元素复制过去。

- 切片的复制：

  ​	使用`copy`函数来复制切片。语法为`copy(destSlice, sourceSlice)`，其中`destSlice`是目标切片，`sourceSlice`是源切片。例如，`s1 := make([]int, 3)`，`s2 := make([]int, 5)`，`copy(s2, s1)`会将`s1`中的元素复制到`s2`中，**最多复制`s1`长度个元素**。在这个例子中，`s2`的前 3 个元素会被`s1`的元素覆盖。

### <5.>map

1. **Map 的创建**

   - **使用`make`函数创建：**

     ​	语法为`make(map[KeyType]ValueType)`，其中`KeyType`是键的类型，`ValueType`是值的类型。例如，`m := make(map[string]int)`创建了一个键为字符串类型、值为整数类型的`map`。这个`map`最初是空的，你可以通过后续操作添加键 - 值对。

   - **字面量创建：**

     ​	可以使用`map`字面量来创建`map`，例如`m := map[string]int{"key1": 1, "key2": 2}`。这种方式在创建`map`的**同时初始化了一些键 - 值对**。在这个例子中，键`"key1"`对应的值是 1，键`"key2"`对应的值是 2。

**Map 的操作**

- 插入和更新键 - 值对：

  ​	例如，对于`m := make(map[string]int)`，可以通过`m["key"] = 3`来插入一个键 - 值对，其中键是`"key"`，值是 3。如果键`"key"`已经存在于`map`中，那么这个操作会更新该键对应的现有值。

- 查找键 - 值对：

  ​	通过键来查找值，例如对于`m := make(map[string]int)`，如果已经插入了键 - 值对`m["key"] = 3`，那么可以通过`value, exists := m["key"]`来查找。其中`value`是键`"key"`对应的实际值，`exists`是一个布尔值，表示键是否存在于`map`中。如果键不存在，`value`会是值类型的零值（对于`int`类型是 0），`exists`为`false`。

  **若访问`m["unknow"]`,不存在则会返回value值类型的0值**。

- 删除键 - 值对：

  ​	使用`delete`函数来删除`map`中的键 - 值对，语法为`delete(mapVariable, key)`。例如，对于`m := make(map[string]int)`，如果已经插入了`m["key"] = 3`，可以通过`delete(m, "key")`来删除这个键 - 值对。

### <6.>函数

**参数传递**：

- **值传递**：在 Go 语言中，默认情况下参数是通过值传递的。这意味着在函数内部对参数的修改不会影响到函数外部的原始值。
- **指针传递**：如果想要在函数内部修改函数外部变量的值，可以通过传递指针的方式。

**返回值：**

​	Go 语言的函数可以有多个返回值。例如：

```go
func calculate(num1, num2 int) (int, int) {
    sum := num1 + num2
    diff := num1 - num2
    return sum, diff
}

result1, result2 := calculate(5, 3)
fmt.Println(result1, result2) // 输出8 2
```

### <7.>type关键字

**定义类型别名**

​	在 Go 语言中，`type`关键字用于定义新的类型别名。这可以让你为已有的类型创建一个新的名字。例如：

```go
type MyInt int
```

**定义结构体类型**

​	如前面提到的结构体部分，`type`关键字与`struct`一起用于定义结构体类型。例如：

```go
type Point struct {
    X int
    Y int
}
```

**定义接口类型**

​	`type`关键字也用于定义接口类型。接口是一种抽象类型，它定义了一组方法签名，但不包含方法的具体实现。例如：

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

**定义函数类型**

​	`type`关键字还可以用于定义函数类型。例如：

```go
type MathOp func(int, int) int
```

​	这里定义了一个名为`MathOp`的函数类型，它表示一个接受两个整数参数并返回一个整数结果的函数。可以像下面这样使用这个函数类型：

```go
func Add(a, b int) int {
    return a + b
}

func main() {
    var op MathOp = Add
    result := op(3, 5)
    fmt.Println(result)
}
```

### <8.>结构体方法

**定义**

​	在 Go 语言中，结构体方法是一种与结构体类型相关联的函数。它可以访问和操作结构体的成员变量，就好像这个方法是结构体本身的一个行为一样。方法的定义使用`func`关键字，并且在函数名和参数列表之间需要指定**接收者**（Receiver）。

**接收者**

​	接收者是一个特殊的参数，它指定了这个方法所属的结构体类型。接收者可以是结构体类型的值或者**指针**。（是否修改内部成员变量）

**值接收者**

```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

指针接收者

```go
func (r *Rectangle) SetWidth(newWidth float64) {
    r.Width = newWidth
}
```

### <9.>go语言输入



```go
	// 这行代码创建了一个从标准输入读取数据的缓冲读取器对象reader。
	reader := bufio.NewReader(os.Stdin)
	// 这里的参数为读取字符串的分隔符	
	// 读到\n的时候截断
	// 系统读入回车发送的是'\r\n'
	// 在后续要对input进行处理来截断'\r\n'
	input, err := reader.ReadString('\n')
	// 清楚尾部的\r\n
	process_input := strings.Trimsuffix(input, '\r\n')
	// 转成想要的数据
	num, err := strings.Atoi(process_input)
```

## 二、实战小项目

### <1.>猜谜游戏

[guessing-game](go-project/01-guessing-game/main.go)

**重点：**go的输入处理。

### <2.>简易在线词典

[simpledict](go-project\02-simpledict\main.go)

**重点：**cURL代码生成、JSON转结构体、JSON分析、网络报文数据格式分析、http报文格式。

### <3.>简易SOCKS5代理

[proxy](go-project\03-proxy\v2\main.go)

**重点：**服务器工作原理、协议报文格式、代理服务器工作原理、并发同步编程（利用goroutine与上下文机制上锁。

## 三、go语言工程实践

### <1.>go语言进阶（并发编程）

#### 1.1 goroutine

**线程与协程：**

|            | 处理器状态 | 调度           | 资源消耗 |
| ---------- | ---------- | -------------- | -------- |
| **协程：** | 用户态     | 由go来调度     | KB       |
| **线程：** | 内核态     | 由操作系统调度 | MB       |

​	线程可以并发的执行多个协程，协程是一种用户态的轻量级线程，它的资源消耗小，可以在单个操作系统线程内并发执行多个任务。

**goroutine特点：**

- goroutine是异步执行的。
- goroutine与匿名函数的组合。
- 用goroutine跑同一个任务，加快速度，但是要注意同步。

#### 1.2 CSP（协程间通讯）

**CSP特点：**

- go提倡**通过通信（通道）来共享内存**而不是通过共享内存实现通信。

#### 1.3 Channel

**语法：**

`make(chan 元素类型, [缓冲大小])`

有缓冲通道：`make(chan int)`

无缓冲通道：`make(chan int, 2)`

**特点：**

- **无缓冲通道**：无缓冲通道在发送数据时会阻塞，直到有接收方准备接收数据。这会导致发送方和接收方必须同时准备好才能进行通信，增加了同步的开销。
- **带缓冲通道**：带缓冲通道在发送数据时不会立即阻塞，而是将数据放入缓冲区。只要缓冲区有空间，发送方就可以继续执行，而不需要等待接收方准备好。这可以显著提高性能，特别是在高并发场景中。

#### 1.4 并发安全Lock

**特点：**

- go仍然支持使用共享内存来通信。

**锁的类型：**

- 互斥锁(sync.Mutex)：确保同一时间只有一个goroutine可以访问该资源。
- 读写锁(sync.RWMutex)：允许多个读者同时访问共享资源，但只允许一个写者访问。
- 条件变量(sync.Cond)：满足特定条件时唤醒等待的goroutine。通常与互斥锁一起使用。
- 原子操作(sync.atomic)：提供无锁的方式来更新共享变量，适用简单的技术和标志位操作。

#### 1.5 WaitGroup（同步goroutine们）

**主要提供的方法：**

1.**`Add(delta int)`**：

- 增加 `WaitGroup` 的计数器。
- 通常在启动 `goroutine` 之前调用，传入需要等待的 `goroutine` 数量。
- 注意：`delta` 必须是正数或零，否则会引发 panic。

2.**`Done()`**：

- 减少 `WaitGroup` 的计数器。
- 通常在 `goroutine` 完成任务后调用。
- 当计数器减到零时，所有调用 `Wait` 的 `goroutine` 会被释放。

3.**`Wait()`**：

- 阻塞当前 `goroutine`，直到 `WaitGroup` 的计数器变为零。
- 通常在主 `goroutine` 中调用，确保所有 `goroutine` 完成后再继续执行。

**使用：**

1. 声明一个WaitGroup，调用Add()传入goroutine的数量。
2. goroutine执行完毕，调用Done()使得计数器减一。
3. 在主gouroutine阻塞，调用Wait()等待所有goroutine执行完毕。

### <2.>go依赖管理

**2.1.1 GOPATH**

- bin：编译的二进制文件
- pkg：项目编译的中间产物
- src： 项目源码

**特点：**

- 无法实现package的多版本控制

**2.1.2 Go Module**

- 通过go.mod文件管理依赖包版本。
- 通过go get/go mod指令工具管理依赖包。
- 实现定义版本规则，通过工具管理项目依赖关系。

[gomod](go-dependencyManagement\go.mod)

#### <3.> 包的管理

gopath方式已经变成过去，去看gomod吧。
[go-package](go-package\src\main\main.go)

## 四、go测试

### <1.>单元测试

​	单元测试用来测试包或者程序的一部分代码或者一组代码的函数。测试单元可以包括**函数、模块**等等。

#### 1.1 单元测试规则

- 所有测试文件以_test.go结尾
- func TestXxx(*testing.T)
- go test -v命令要在xx_test.go文件目录下执行，如果在后面指定某个包，要使用相对路径或者是绝对路径
- 可以用第三方包来帮助测试---github.com/stretchr/testify/assert
- 初始化逻辑放到TestMain中

```go
func TestMain(m *testing.T){
	// 测试前：数据装载、配置初始化等前置工作
	code := m.Run()
	// 测试后：释放资源等收尾工作2
	os.Exit(code)
}
```

#### 1.2 单元测试-覆盖率

**go命令：**

​	go test judgePassLine_test.go judgePassLine.go --cover

**覆盖率：**

​	单元测试覆盖率是指单元测试覆盖的**代码量占总代码量的比例**。

#### 1.3 单元测试-Tips

- 一般覆盖率：50%~60%， 较高覆盖率80%。
- 测试分支相互独立、全面覆盖。
- 测试单元粒度足够小，函数单一职责（设计模式）。

#### 1.5 单元测试-依赖

外部依赖 => 稳定（单元测试能任何时间任何函数独立运行）&幂等（每次case运行结果一样）

#### 1.6 单元测试-mock

- 用一些方法来模仿测试代码所需要的外部资源
- 为一个函数打桩（）
- 为一个方法打桩（）
- 在软件测试中，"为函数打桩"（也称为“桩函数”或“mocking”）是一种技术，用于模拟或替换被测试代码中依赖的外部组件或函数。这样做可以隔离被测试的代码，使其不受外部依赖的影响，从而更容易、更可靠地进行单元测试。
- 第三方包来mock --bou.ke/monkey
- [mock-test](go-test/fileprocess/readFile_test.go)

### <2.> 基准测试

**功能：**

​	基准测试（Benchmark Testing）是一种性能测试方法，用于测量代码在特定条件下的执行速度和效率。基准测试可以帮助你了解代码的性能瓶颈，优化代码，并比较不同实现的性能差异。

[go-test指令问题](go-test/serverindex/serverindex_test.go)

[go_test指令问题解决]([Golang 单元测试、基准测试、并发基准测试 - 林锅 - 博客园](https://www.cnblogs.com/linguoguo/p/10371253.html))



