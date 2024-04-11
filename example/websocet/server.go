package wsexample

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	var conn, err = upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error: ", err)
		return
	}
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %v %s", mt, message)

		// 返回同样的消息给客户端
		err = conn.WriteMessage(mt, []byte(string(message)+"吗？"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	log.Println("服务关闭!")
	conn.Close()
}

func ExecWSServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8888", nil)
}
