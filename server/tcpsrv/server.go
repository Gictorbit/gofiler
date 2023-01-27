package tcpsrv

import (
	"github.com/Gictorbit/gofiler/proto/pb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"net"
	"sync"
)

const (
	PacketMaxByteLength = 2048
	ServerSocketType    = "tcp"
)

type Empty struct{}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitChan   chan Empty
	wg         sync.WaitGroup
	log        *zap.Logger
}

func NewServer(listenAddr string, logger *zap.Logger) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitChan:   make(chan Empty),
		wg:         sync.WaitGroup{},
		log:        logger,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen(ServerSocketType, s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptConnections()

	<-s.quitChan
	return nil
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			s.log.Error("accept connection error", zap.Error(err))
			continue
		}
		s.log.Info("new Connection to the server", zap.String("Address", conn.RemoteAddr().String()))
		go s.HandleConnection(conn)
	}
}

func (s *Server) HandleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		packet, err := s.ReadPacket(conn)
		if err != nil {
			s.log.Error("read packet error", zap.Error(err))
			continue
		}
		switch packet.MessageType {
		case pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_REQUEST:
			req := &pb.UploadFileRequest{}
			if e := proto.Unmarshal(packet.Payload, req); e != nil {
				s.log.Error("unmarshal upload request failed", zap.Error(err))
				continue
			}
			if e := s.UploadFileHandler(req, conn); e != nil {
				s.log.Error("handle upload file failed", zap.Error(err))
				continue
			}
		}
	}
}

func (s *Server) Stop() {
	s.wg.Wait()
	s.quitChan <- Empty{}
}
