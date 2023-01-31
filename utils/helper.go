package utils

import (
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/Gictorbit/gofiler/proto/pb"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"os"
	"path/filepath"
)

const (
	PacketMaxByteLength = 2048
	ServerSocketType    = "tcp"
)

var (
	ErrInvalidPacketSize = errors.New("invalid packet size")
)

func FileInfo(fPath string, file *os.File) (*pb.File, error) {
	fileExtension := filepath.Ext(fPath)
	fileStat, err := os.Stat(fPath)
	if err != nil {
		return nil, err
	}
	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return nil, err
	}
	return &pb.File{
		Name:      file.Name(),
		Size:      fileStat.Size(),
		Extension: fileExtension,
		Checksum:  fmt.Sprintf("%x", hash.Sum(nil)),
	}, nil
}

type MessageBody[T proto.Message] struct {
	MessageType pb.MessageType
	Payload     []byte
	Message     T
}

func ReadMessageFromConn[T proto.Message](conn net.Conn, message T) (*MessageBody[T], error) {
	buf := make([]byte, PacketMaxByteLength)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	packet := &MessageBody[T]{
		Payload:     buf[1:n],
		MessageType: pb.MessageType(buf[0]),
	}
	if e := proto.Unmarshal(packet.Payload, message); e != nil {
		return nil, e
	}
	packet.Message = message

	return packet, nil
}

func GenerateRandomCode(codeLength int) string {
	b := make([]byte, codeLength)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}
