package internal

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// 客户端连接状态列表
	clients map[*Client]bool

	// 来自客户端的入站消息
	broadcast chan []byte

	// 客户端注册队列
	register chan *Client

	// 客户端注销队列
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {

		// 读取注册队列，更新客户端状态列表
		case client := <-h.register:
			h.clients[client] = true

		// 读取注销队列，从状态列表中删除客户端信息，并关闭客户端连接
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		// 读取广播队列中的消息，将消息写入每一个客户的发送队列
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		}
	}
}
