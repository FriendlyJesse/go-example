package wsexample

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) readPump() {
	// ws 结束后注销当前 client 和关闭 conn
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	// 读取大小
	c.conn.SetReadLimit(maxMessageSize)
	// 如果在该期限内没有从连接中读取到数据，读取操作将返回超时错误。这可以帮助防止程序在长时间等待数据时被阻塞。
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// 在接收到Pong消息时更新读取截止时间，默认情况下 Chrome 会自动发送 Ping 消息
	c.conn.SetPongHandler(func(string) error {
		fmt.Println(1010101010)
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		mt, message, err := c.conn.ReadMessage()
		fmt.Println("mt: ", mt)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			// 为什么它能知道客户端断开了连接
			log.Printf("error: %v", err) // error: websocket: close 1001 (going away)
			break
		}
		// 去除消息汇总的空格和回车
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// 广播消息
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	// 定时器
	ticker := time.NewTicker(pingPeriod)
	// ws 结束后停止 ticker 和关闭 conn
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send: // 收到消息后广播消息
			// 如果在该期限内无法将数据写入连接，写入操作将返回超时错误。这有助于确保写入操作在合理的时间内完成，避免程序长时间等待。
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			fmt.Println(c.hub.clients)
			fmt.Println("收到的消息：", string(message))

			// 发送消息
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			// 超时时间
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
