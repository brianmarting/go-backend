package socket

import (
	"go-backend/app/socket/tcp"
)

func NewTcpSocketListener() tcp.Listener {
	return tcp.NewListener()
}
