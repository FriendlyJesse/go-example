package example

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func tcpPipe(conn *net.TCPConn) {
	// TCP连接的地址
	var ipStr = conn.RemoteAddr().String()
	// 关闭连接
	defer func() {
		fmt.Printf("%v 失去连接\n", ipStr)
		conn.Close()
	}()

	// 获取TCP连接对象的数据流
	var reader = bufio.NewReader(conn)
	// 接收并返回消息
	for {
		// 获取接收数据
		message, err := reader.ReadString('\n')
		// 连接异常
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println("服务端接收数据：", message)
		time.Sleep(3 * time.Second)

		// 发送数据
		msg := conn.RemoteAddr().String() + "--服务端发送数据\n"
		b := []byte(msg)
		conn.Write(b)
	}
}

func tcpServer() {
	// TCP 对象
	var tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	// 创建 TCP 监听对象
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()

	for {
		// 通过 TCP 监听对象获取与客户端的TCP连接对象
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 连接成功后创建协程去处理
		go tcpPipe(tcpConn)
	}
}

func onMessageReceived(conn *net.TCPConn) {
	// 创建TCP连接对象的IO
	var reader = bufio.NewReader(conn)
	// 发送数据
	var b = []byte(conn.LocalAddr().String() + "客户端在发送数据。\n")
	conn.Write(b)

	for {
		// 获取 TCP连接对象的数据流
		var msg, err = reader.ReadString('\n')
		fmt.Println("客户端收到服务端数据：", msg)
		if err != nil || err == io.EOF {
			break
		}
		time.Sleep(2 * time.Second)
		// 通过TCP连接对象发送数据给服务端
		_, err = conn.Write(b)
		if err != nil {
			break
		}
	}
}

func tcpClient() {
	var tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("客户端连接错误：", err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("客户端连接成功：", conn.LocalAddr().String())
	onMessageReceived(conn)
}

func ExecTCP() {
	// 启用两个 shell 执行代码
	// 启用 server
	tcpServer()

	// 启用 client
	tcpClient()
}
