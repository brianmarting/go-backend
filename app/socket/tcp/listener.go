package tcp

import (
	"bufio"
	"fmt"
	"go-backend/app/socket"
	"net"
	"os"

	"github.com/rs/zerolog/log"
)

type Listener struct {
	port string
}

func NewListener() Listener {
	port := os.Getenv("TCP_PORT")

	return Listener{
		port: port,
	}
}

func (l Listener) Start() (<-chan socket.Message, chan<- string, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", l.port))
	defer listener.Close()
	if err != nil {
		return nil, nil, err
	}

	conn, err := listener.Accept()
	if err != nil {
		return nil, nil, err
	}

	log.Info().Msg(fmt.Sprintf("started listening on tcp port %s", l.port))

	inboundMessageChannel := make(chan socket.Message)
	outboundMessageChannel := make(chan string)

	go func() {
		defer close(inboundMessageChannel)
		for {
			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Info().Err(err).Msg("something went wrong when reading tcp msg")
				continue
			}

			inboundMessageChannel <- convertToMessage(msg)

			resp := <-outboundMessageChannel
			if _, err := conn.Write([]byte(resp)); err != nil {
				log.Info().Err(err).Msg("something went wrong when sending tcp msg reply")
			}
		}
	}()

	return inboundMessageChannel, outboundMessageChannel, nil
}

func convertToMessage(message string) socket.Message {
	return Message{
		body: []byte(message),
	}
}
