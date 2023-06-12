package socket

import (
	"go-backend/internal/app/socket"
	"go-backend/internal/app/socket/tcp"
)

func NewTcpSocketListener(port string) socket.Listener {
	return tcp.NewListener(port)
}
