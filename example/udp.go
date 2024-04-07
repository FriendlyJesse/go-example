package example

import (
	"fmt"
	"io"
	"net"
	"time"
)

func udpServer() {
	var udpAddr, _ = net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	for {
		// 接收数据
		var message = make([]byte, 4096)
		// n：数据长度; addr：客户端 IP 地址
		var n, addr, err = conn.ReadFromUDP(message)
		if err != nil || err == io.EOF {
			break
		}

		fmt.Println("服务端接收数据：", string(message[:n]))
		time.Sleep(3 * time.Second)

		// 发送数据
		msg := conn.LocalAddr().String() + "--服务端发送数据"
		b := []byte(msg)
		conn.WriteToUDP(b, addr)
	}
}

func udpClient() {
	var udpAddr, _ = net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	// 创建 udp 连接对象
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("客户端连接错误：", err.Error())
		return
	}
	fmt.Println("客户端连接成功：", conn.LocalAddr().String())

	for {
		var b = []byte(conn.LocalAddr().String() + "--客户端发送数据")
		// 发送数据
		conn.Write(b)
		time.Sleep(2 * time.Second)

		message := make([]byte, 4096)
		n, _, err := conn.ReadFromUDP(message)
		fmt.Println("客户端收到数据：", string(message[:n]))
		if err != nil || err == io.EOF {
			break
		}

	}

}

func ExecUDP() {
	// server
	udpServer()

	// client
	udpClient()
}
