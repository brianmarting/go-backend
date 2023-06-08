package socket

import (
	"go-backend/app/socket/tcp"
)

func NewTcpSocketListener(port string) tcp.Listener {
	return tcp.NewListener(port)
}
