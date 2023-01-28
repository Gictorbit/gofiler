package tcpclient

import (
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
	"google.golang.org/protobuf/proto"
	"net"
)

type PacketBody struct {
	MessageType pb.MessageType
	Payload     []byte
}

func (c *Client) ReadPacket(conn net.Conn) (*PacketBody, error) {
	buf := make([]byte, utils.PacketMaxByteLength)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return &PacketBody{
		MessageType: pb.MessageType(buf[0]),
		Payload:     buf[1:n],
	}, nil
}

func (c *Client) SendRequestPacket(packet *PacketBody) error {
	buf := make([]byte, 0)
	buf = append(buf, byte(packet.MessageType))
	buf = append(buf, packet.Payload...)
	if len(buf) > utils.PacketMaxByteLength {
		return utils.ErrInvalidPacketSize
	}
	if _, err := c.conn.Write(buf); err != nil {
		return err
	}
	return nil
}

func (c *Client) SendRequest(msgType pb.MessageType, msg proto.Message) error {
	respBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	packet := &PacketBody{
		MessageType: msgType,
		Payload:     respBytes,
	}
	if e := c.SendRequestPacket(packet); e != nil {
		return e
	}
	return nil
}
