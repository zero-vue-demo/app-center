package internal

import (
	"bytes"
	"fmt"
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

	// send buffer size
	bufSize = 256
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// 从 websocket 中读取客户端消息
//
// 应用程序在每个连接的 Goroutine 中运行 readPump。
// 该应用程序通过执行来自该 Goroutine 的 所有读取来确保一个连接上最多有一个读取器。
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	// 限制读取客户端的消息大小
	c.conn.SetReadLimit(maxMessageSize)
	// 设置读取的超时时间
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// 设置接收到客户端心跳反馈时的处理程序
	c.conn.SetPongHandler(func(string) error {
		// 重新设置读取超时时间
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		// 从 websocket 中读取客户端消息
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v", err)
			}
			break
		}
		// 格式化接收到的消息
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// 将消息写入广播队列
		c.hub.broadcast <- message
	}
}

// 向客户端发送消息
//
// 为每个连接启动一个运行 writePump 的 Goroutine。
// 该应用程序通过执行来自该 Goroutine 的 所有写操作来确保一个连接最多有一个编写器。
func (c *Client) writePump() {
	// 创建一个周期性的定时器
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {

		// 心跳周期触发
		case <-ticker.C:

			// 设置 10s 写超时
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			// 发送心跳包，ping 客户端
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		// 消息队列触发
		case message, ok := <-c.send:

			// 设置 10s 写超时
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			// 若消息获取错误
			if !ok {
				// 发送关闭信号
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				// 关闭协程
				return
			}

			// 发送刚从队列中读取到的一条文本消息
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// 若队列未清空，则继续发送，直到搬空队列
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			// 关闭写通道
			if err := w.Close(); err != nil {
				return
			}

		}
	}
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, bufSize),
	}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
