// socks5代理-TCP echo server
// 代理服务器的本质就是服务器上安装了一个代理软件，这个软件具有转发包的功能。
/* TCP echo server（传输控制协议回显服务器）是一种基于 TCP 协议的服务器程序。
 * 它的主要功能是接收客户端发送过来的 TCP 数据，并将接收到的数据原封不动地发送（回显）回客户端。
 * 这种服务器通常用于网络测试、调试工具以及一些简单的网络通信实验。*/
package main

// 这是个服务器，把它运行起来需要在另一个地方建立连接使用

import (
	"bufio"
	//"bytes"
	//"go/build"
	"log"
	"net"
)

func main() {
	// 基于TCP监听127.0.0.1:1080
	// 如果监听成功，返回一个net.Listener的对象
	// Listener对象用于后续接受客户端的连接请求
	// 创建了一个TCP服务程序
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}

	for {
		/* 在服务器端阻塞等待客户端的连接请求，
		 * 并在接收到一个新的客户端连接时进行相应的处理。*/
		/* 返回一个net.Conn对象，这个对象代表了服务器与该客户端的连接。
		 * 通过这个对象，服务器可以与客户端进行双向的数据读写操作。*/
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accepy failed %v", err)
			continue
		}
		/* 用于启动一个新的goroutine。
		 * goroutine是 Go 语言特有的轻量级线程，可以在程序中实现并发执行。
		 * 与操作系统线程相比，goroutine的开销非常小，创建和销毁goroutine的成本很低
		 * 这允许服务器可以并发的处理多个客户端的请求。*/
		go process(client)
	}

}

func process(conn net.Conn) {
	defer conn.Close()

	// 创建一个带缓冲的流
	// 带缓冲的流可以减少底层系统的调用次数
	// 通过网络连接读取数据
	reader := bufio.NewReader(conn)

	// 这里是一个死循环。
	for {
		// 读取一个字节的数据
		b, err := reader.ReadByte()
		if err != nil {
			break
		}

		// 将该字节写如conn,返回客户端
		// 将b转换成一个字节切片，方便处理不同长度的数据写入
		// _代表写入的字节数，这里忽略
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
	// tag_str := " - Meteora"
	
	// /* append函数用于向切片中追加元素。
	//  * 它是一个可变参数函数，可以接受一个切片和零个或多个值，并返回一个新的切片。
	//  * 这个新切片包含了原始切片的所有元素以及新添加的元素。 */

	// _, err := conn.Write([]byte(tag_str))
	// if err != nil {
	// 	panic(err)
	// }

}
