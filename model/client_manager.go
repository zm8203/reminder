package model

import "github.com/gorilla/websocket"

// Client 管理 WebSocket 连接
type Client struct {
	conn         *websocket.Conn
	reminderChan chan Reminder
}

var clients []*Client

// NewClient 创建新的客户端
func NewClient(conn *websocket.Conn, reminderChan chan Reminder) *Client {
	return &Client{
		conn:         conn,
		reminderChan: reminderChan,
	}
}

// AddClient 添加客户端到列表中
func AddClient(client *Client) {
	clients = append(clients, client)
}

// RemoveClient 从列表中移除客户端
func RemoveClient(client *Client) {
	for i, c := range clients {
		if c == client {
			clients = append(clients[:i], clients[i+1:]...)
			return
		}
	}
}

// BroadcastReminder 广播提醒给所有客户端
func BroadcastReminder(reminder Reminder) {
	for _, client := range clients {
		client.reminderChan <- reminder
	}
}
