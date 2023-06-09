package socket

import (
	"go-backend/internal/app/socket/tcp"
)

func NewTcpSocketListener(port string) tcp.Listener {
	return tcp.NewListener(port)
}
