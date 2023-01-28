package tcpsrv

import (
	"errors"
	"github.com/Gictorbit/gofiler/proto/pb"
	"google.golang.org/protobuf/proto"
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
	buf := make([]byte, PacketMaxByteLength)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return &PacketBody{
		MessageType: pb.MessageType(buf[0]),
		Payload:     buf[1:n],
	}, nil
}

func (s *Server) SendResponsePacket(conn net.Conn, packet *PacketBody) error {
	buf := make([]byte, 0)
	buf = append(buf, byte(packet.MessageType))
	buf = append(buf, packet.Payload...)
	if len(buf) > PacketMaxByteLength {
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
