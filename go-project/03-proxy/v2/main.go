// socks5服务器
/* 1.认证支持。SOCKS5支持多种认证方式，包括无需认证，使用用户名和密码认证等多种方式，
 *	 客户端和代理服务器可以协商选择合适的认证方法。
 * 2.通信协议无关性。SOCKS5协议可以应用于各种应用层协议的代理，如HTTP、SMTP、FTP等。
 * 3.UDP支持。SOCKS5支持用户数据报（UDP）的代理，像是需要实时传输的音视频相关应用。 */

package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

// 协议名字以及版本
// 可以隐式的当作byte来比较
const socks5Ver = 0x05
const cmdBind = 0x01
const atypIPV4 = 0x01
const atypHOST = 0x03
const atypIPV6 = 0x04

/*
运行时，先运行代理服务器程序，然后在客户端通过代理服务器向其他服务器通信
* 在客户端运行curl --socks5 127.0.0.1:1080 -v http://www.qq.com
* curl会根据 SOCKS5 协议的默认行为来确定ver、nmethods和method的值。
* 这里所说的SOCKS5是标准的SOCKS5，而后面接受报文处理的是这里写的SOCKS5
* <1>.curl会将协议版本设置为0x05，这是 SOCKS5 协议的标准版本号，
* 表示它正在尝试使用 SOCKS5 协议与代理服务器进行通信
* <2>.默认情况下，curl会发送两种认证方法，即0x00（无需认证）和0x02（用户名 / 密码认证）。
* 所以nmethods的值为2，method的值为[0x00, 0x02]。
*/
func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}

		go process(client)
	}

}

func process(client net.Conn) {
	defer client.Close()
	reader := bufio.NewReader(client)

	// 用auth函数处理这个客户端连接，用reader来解析客户端的报文
	err := auth(reader, client)
	if err != nil {
		log.Panicf("client %v auth failed: %v", client.RemoteAddr(), err)
		return
	}
	log.Println("auth success")

	err = connect(reader, client)
	if err != nil {
		log.Panicf("client %v connect failed: %v", client.RemoteAddr(), err)
		return
	}
}

// 第一阶段，协商
/* 客户端需要先和代理服务器建立连接，客户端通过代理服务器的认证。
 * 客户端需要发送协商报文，之后代理服务器返回认证报文
 */
func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// 客户端发送的报文
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  | version, number of methods, methods
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 | 协议版本， 客户端支持的方法数量， 方法列表
	// +----+----------+----------+
	// 对于methods,0x00：无需认证， 0x02：用户名/密码认证

	// 读取第一个字节: 协议版本号，如果是SOCKS5的版本号，则继续处理
	ver, err := reader.ReadByte()
	if err != nil {
		// fmt.Errorf可以创建错误对--包含格式化字符串的err对象
		return fmt.Errorf("read ver failed: %v", err)
	}
	// go中字节类型其实是一个无符号的8位整数
	if ver != socks5Ver {
		return fmt.Errorf("not supported: %v", err)
	}

	// 读取第二个字节：客户端支持的认证方法数量
	methodSeize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read nmethods failed: %v", err)
	}

	// 读取“剩下的”字节：客户端支持的协议列表
	// 初始化一个字节切片，大小为methodSize，用来存协议列表
	method := make([]byte, methodSeize)
	// 使用io.readFull从reader中读取内容填充到method
	/* io.ReadFull函数有一个特性，第二个参数有多大，
	 * 就从reader中读取对应的字节，填充满第二个参数。*/
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read methods failed: %v", err)
	}
	// 成功后，打印日志，包括版本号以及客户端支持的每种认证方法
	// 打印字节切片时，它会将字节切片中的字节以十进制数字的形式打印出来，字节之间用空格分隔。
	log.Println("ver: ", ver, ",method: ", method)

	// 选择认证方式并返回报文
	/* 服务器端对协商报文解析后，需要在客户端支持的认证方式中选择一种，并返回报文。
	 * conn代表此前客户端与服务器建立的连接。使用Write函数写入字节流，
	 * 内容包括协议版本号和选中的认证方法。Write函数像客户端返回的报文
	 */
	// +----+----------+
	// |VER | METHOD   |
	// +----+----------+
	// | 1  |    1     |
	// +----+----------+
	// 服务器选择无需认证的形式：0x00
	/* curl的输出主要显示了它自身的操作过程和遇到的问题，
	 * 但默认情况下它不会直接显示从代理服务器接收到的详细字节流响应报文。*/
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write failed%v", err)
	}

	return nil
}

// 第二阶段：请求阶段，客户端向代理服务器发送请求，指出要通过代理访问的IP地址或域名。
// 第三阶段，relay阶段，与最终服务器建立TCP连接。
func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 	版本号，socks5的值为0x05
	// CMD 	命令类型，0x01表示CONNECT请求
	// RSV 	未来扩展，保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名，第一个字节是域名长度
	//   0x02表示IPv6地址
	// DST.ADDR 目标地址，一个可变长度的值
	// DST.PORT 目标端口，固定2个字节，大端序

	// 先校验前几个字段的值
	// 4个字节也正好是后续IPv4的长度
	buff := make([]byte, 4)
	// 这里已经读取了四个字节
	_, err = io.ReadFull(reader, buff)
	ver, cmd, atyp := buff[0], buff[1], buff[3]
	if ver != socks5Ver {
		// %w可以添加更多上下文信息，定位到是那个阶段出了问题
		return fmt.Errorf("read header failed:%w", err)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supported cmd: %v", err)
	}

	// 目标地址
	addr := ""
	switch atyp {
	case atypIPV4:
		// 再读四个刚好是IPv4地址
		_, err = io.ReadFull(reader, buff)
		if nil != err {
			return fmt.Errorf("read atype failed: %w", err)
		}
		// fmt.Sprintf()格式化字符串返回，返回为点分十进制格式
		addr = fmt.Sprintf("%d.%d.%d.%d", buff[0], buff[1], buff[2], buff[3])
	case atypHOST:
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read hostSize failed: %w", err)
		}

		// 读入域名，转为字符串
		host := make([]byte, hostSize)
		_, err = io.ReadFull(reader, host)
		addr = string(host)
	case atypIPV6:
		// 这里用errors.New(), 因为没有额外的上下文信息
		return errors.New("IPv6: not supported yet")
	default:
		return errors.New("invalid atype")
	}

	// 读取最后的端口号，直接复用buff容器
	_, err = io.ReadFull(reader, buff[:2])
	if err != nil {
		return fmt.Errorf("read port failed: %w", err)
	}
	// 使用大端序来解析端口号
	port := binary.BigEndian.Uint16(buff[:2])

	log.Println("analysis--dial", addr, port)

	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP 操作的结果，Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// 通过下列的地址与目标地址建立的连接
	// ATYPE 地址类型
	// BND.ADDR 服务绑定的地址，下列假设为四个字节设为0
	// BND.PORT 服务绑定的端口DST.PORT，两个字节
	// 回复的报文，共10字节
	// curl指令默认不会输出从代理服务器返回的内容
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}

/* 第三阶段，（代理服务器）与目标服务器建立tcp连接------------------------ */
	// Dial()函数用来建立连接，返回一个net.Conn的对象
	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil{
		return fmt.Errorf("dial dst failed:%w", err)
	}

	defer dest.Close()
	log.Println("connect--dial", addr, port)

	// 实现数据的双向传输
	// goroutine是轻量级线程，在这里并发执行
	/* 建立两个goroutine，分别为reader到dest、dest到reader。
	 * 但是在这里2个goroutine与主进程同时运行，如果主进程运行的更快，
	 * 此时双向数据可能还没有完成，该函数就关闭了。
	 * 上下文控制ctx，不仅为了防止主程序提前结束，而且还会协调goroutine之间的操作
	 * 拿到一个上下文空间，把下列运行缩进去控制
	*/
	// “上锁”， “钥匙：cancel()”
	/* io.Copy与context的协同工作机制
	 * io.Copy函数在执行过程中会检查context的状态。
	 * 当ctx被取消（ctx.Done()通道被关闭），io.Copy会停止数据的复制操作，释放相关资源。
	 * 这里只要有一个goroutine执行了cancel()，那么ctx就会关闭，传输数据也会关闭*/
	ctx, cancel := context.WithCancel(context.Background())
	// 若没有进行传输数据，保证ctx被取消
	defer cancel()

	// io.Copy()函数主要实现了从io.Reader类型的对象复制到一个io.Writer的对象
	// net.Conn同时实现了io.Reader和io.Writer接口。
	go func() {
		// 客户端到服务端
		_, _ = io.Copy(dest, reader)
		// 说明该goroutine已经结束
		cancel()
	}()

	go func() {
		// 服务端到客户端
		_, _ = io.Copy(conn, dest)
		cancel()
	}()
	
	// 阻塞主程序，使得主程序等待双向传输结束
	// ctx.Done()返回的是一个通道类型， <-表示它正在接收信号
	// 等调用cancel()解除ctx，他就会关闭
	<-ctx.Done()

	return nil
}
