package tcpsrv

import (
	"errors"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/server/storage"
	"go.uber.org/zap"
	"net"
)

// FileInfoHandler handles file info request returns all information about given file share code
func (s *Server) FileInfoHandler(req *pb.FileInfoRequest, conn net.Conn) error {
	fileInfo, err := s.fileStorage.FileInfo(req.GetShareCode())
	response := &pb.FileInfoResponse{
		FileInfo:     fileInfo,
		ResponseCode: pb.ResponseCode_RESPONSE_CODE_SUCCESS,
	}
	if err != nil {
		if errors.Is(err, storage.ErrFileNotFound) {
			s.log.Warn("file no found", zap.String("ShareCode", req.GetShareCode()))
			response.ResponseCode = pb.ResponseCode_RESPONSE_CODE_NOT_FOUND
			return s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_FILE_INFO_RESPONSE, response)
		}
		s.log.Error("get file info failed",
			zap.Error(err),
			zap.String("ShareCode", req.GetShareCode()),
		)
		response.ResponseCode = pb.ResponseCode_RESPONSE_CODE_FAILED
		if e := s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_FILE_INFO_RESPONSE, response); e != nil {
			return e
		}
		return err
	}
	return s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_FILE_INFO_RESPONSE, response)
}
