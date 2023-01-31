package tcpsrv

import (
	"errors"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/server/storage"
	"go.uber.org/zap"
	"net"
)

// DownloadFileHandler handles download file request and sends a file to client
func (s *Server) DownloadFileHandler(req *pb.DownLoadFileRequest, conn net.Conn) error {
	fileInfo, err := s.fileStorage.FileInfo(req.GetIdCode())
	response := &pb.DownLoadFileResponse{
		File:         fileInfo,
		ResponseCode: pb.ResponseCode_RESPONSE_CODE_READY,
	}
	if err != nil {
		if errors.Is(err, storage.ErrFileNotFound) {
			s.log.Warn("file no found", zap.String("IDCode", req.IdCode))
			response.ResponseCode = pb.ResponseCode_RESPONSE_CODE_NOT_FOUND
			return s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_DOWNLOAD_FILE_RESPONSE, response)
		}
		s.log.Error("get file info failed",
			zap.Error(err),
			zap.String("IDCode", req.IdCode),
		)
		response.ResponseCode = pb.ResponseCode_RESPONSE_CODE_FAILED
		if e := s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_DOWNLOAD_FILE_RESPONSE, response); e != nil {
			return e
		}
		return err
	}
	if e := s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_DOWNLOAD_FILE_RESPONSE, response); e != nil {
		return e
	}
	if e := s.fileStorage.GetFile(conn, fileInfo); e != nil {
		s.log.Error("error get File",
			zap.Error(err),
			zap.String("FileName", fileInfo.Name),
			zap.Int64("FileSize", fileInfo.Size),
			zap.String("Md5Sum", fileInfo.Checksum),
			zap.String("IDCode", fileInfo.IdCode),
		)
		response.ResponseCode = pb.ResponseCode_RESPONSE_CODE_FAILED
		return s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_DOWNLOAD_FILE_RESPONSE, response)
	}
	s.log.Info("sent file successfully",
		zap.Error(err),
		zap.String("FileName", fileInfo.Name),
		zap.Int64("FileSize", fileInfo.Size),
		zap.String("Md5Sum", fileInfo.Checksum),
		zap.String("IDCode", fileInfo.IdCode),
	)
	response.ResponseCode = pb.ResponseCode_RESPONSE_CODE_SUCCESS
	return s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_DOWNLOAD_FILE_RESPONSE, response)
}
