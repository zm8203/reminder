package api

import (
	"awesomeProject/model"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ServeWebSocket 服务 WebSocket
func ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	// 监听提醒通道
	reminderChan := make(chan model.Reminder)

	// 将新连接添加到客户端管理器中
	client := model.NewClient(conn, reminderChan)
	model.AddClient(client)

	defer func() {
		model.RemoveClient(client)
		conn.Close()
	}()

	// 接收 WebSocket 消息
	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}()

	// 推送提醒消息给客户端
	for {
		select {
		case reminder := <-reminderChan:
			err := conn.WriteJSON(reminder)
			if err != nil {
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
}
