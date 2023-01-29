package tcpsrv

import (
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"net"
	"sync"
)

type Empty struct{}

type Server struct {
	listenAddr string
	storage    string
	ln         net.Listener
	quitChan   chan Empty
	wg         sync.WaitGroup
	log        *zap.Logger
}

func NewServer(listenAddr string, logger *zap.Logger, storagePath string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitChan:   make(chan Empty),
		wg:         sync.WaitGroup{},
		log:        logger,
		storage:    storagePath,
	}
}

func (s *Server) Start() {
	ln, err := net.Listen(utils.ServerSocketType, s.listenAddr)
	if err != nil {
		s.log.Error("failed to listen", zap.Error(err))
		return
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptConnections()
	s.log.Info("server started",
		zap.String("ListenAddress", s.listenAddr),
		zap.String("storagePath", s.storage),
	)
	<-s.quitChan
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
	packet, err := s.ReadPacket(conn)
	if err != nil {
		s.log.Error("read packet error", zap.Error(err))
		return
	}
	switch packet.MessageType {
	case pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_REQUEST:
		req := &pb.UploadFileRequest{}
		if e := proto.Unmarshal(packet.Payload, req); e != nil {
			s.log.Error("unmarshal upload request failed", zap.Error(err))
			return
		}
		if e := s.UploadFileHandler(req, conn); e != nil {
			s.log.Error("handle upload file failed", zap.Error(err))
			return
		}
	}

}

func (s *Server) Stop() {
	s.wg.Wait()
	s.quitChan <- Empty{}
	s.log.Info("stop server")
}
