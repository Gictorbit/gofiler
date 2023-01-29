package tcpsrv

import (
	"github.com/Gictorbit/gofiler/proto/pb"
	"net"
)

func (s *Server) DeleteFileHandler(req *pb.DeleteFileRequest, conn net.Conn) (err error) {
	return nil
}
