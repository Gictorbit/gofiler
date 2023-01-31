package tcpsrv

import (
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
	"go.uber.org/zap"
	"net"
)

func (s *Server) UploadFileHandler(req *pb.UploadFileRequest, conn net.Conn) error {
	req.File.IdCode = utils.GenerateRandomCode(10)
	resultResp := &pb.UploadFileResponse{
		IdCode:       req.File.IdCode,
		ResponseCode: pb.ResponseCode_RESPONSE_CODE_SUCCESS,
	}
	readyResponse := &pb.UploadFileResponse{
		ResponseCode: pb.ResponseCode_RESPONSE_CODE_READY,
	}
	if e := s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, readyResponse); e != nil {
		return e
	}
	if err := s.fileStorage.SaveFile(req.File, conn); err != nil {
		s.log.Error("error saving File",
			zap.Error(err),
			zap.String("FileName", req.GetFile().Name),
			zap.Int64("FileSize", req.GetFile().Size),
			zap.String("Md5Sum", req.GetFile().Checksum),
			zap.String("IDCode", req.GetFile().IdCode),
		)
		resultResp.ResponseCode = pb.ResponseCode_RESPONSE_CODE_FAILED
		return s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, resultResp)
	}
	s.log.Info("received a new file successfully",
		zap.String("FileName", req.GetFile().Name),
		zap.Int64("FileSize", req.GetFile().Size),
		zap.String("Md5Sum", req.GetFile().Checksum),
		zap.String("IDCode", req.GetFile().IdCode),
	)
	return s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, resultResp)
}
