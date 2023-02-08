package tcpsrv

import (
	"errors"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
)

var (
	ErrInvalidPacketSize = errors.New("invalid packet size")
)

type PacketBody struct {
	MessageType pb.MessageType
	Payload     []byte
}

func (s *Server) ReadPacket(conn net.Conn) (*PacketBody, error) {
	buf := make([]byte, utils.PacketMaxByteLength)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, io.EOF
	}
	messageType := pb.MessageType(buf[0])
	payload := buf[1:n]
	return &PacketBody{
		MessageType: messageType,
		Payload:     payload,
	}, nil
}

func (s *Server) SendResponsePacket(conn net.Conn, packet *PacketBody) error {
	buf := make([]byte, 0)
	buf = append(buf, byte(packet.MessageType))
	buf = append(buf, packet.Payload...)
	if len(buf) > utils.PacketMaxByteLength {
		return ErrInvalidPacketSize
	}
	if _, err := conn.Write(buf); err != nil {
		return err
	}
	return nil
}

func (s *Server) SendResponse(conn net.Conn, msgType pb.MessageType, msg proto.Message) error {
	respBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	packet := &PacketBody{
		MessageType: msgType,
		Payload:     respBytes,
	}
	if e := s.SendResponsePacket(conn, packet); e != nil {
		return e
	}
	return nil
}
