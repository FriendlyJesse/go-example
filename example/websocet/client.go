package wsexample

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func send(conn *websocket.Conn) {
	for {
		var reader = bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, []byte(l))
	}
}

func ExecWSClient() {
	dl := websocket.Dialer{}
	conn, _, _ := dl.Dial("ws://127.0.0.1:8888", nil)

	go send(conn)

	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		fmt.Println(m, string(p))
	}
}
