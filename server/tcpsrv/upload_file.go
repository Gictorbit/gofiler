package tcpsrv

import (
	"errors"
	"fmt"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
	"go.uber.org/zap"
	"io"
	"net"
	"os"
	"path"
	"strings"
)

func (s *Server) UploadFileHandler(req *pb.UploadFileRequest, conn net.Conn) (err error) {
	filePath := path.Join(s.storage, strings.TrimSpace(req.File.Name))
	if _, err = os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
		existsResp := &pb.UploadFileResponse{
			ResponseCode: pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_EXISTS,
		}
		if e := s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, existsResp); e != nil {
			return e
		}
		s.log.Warn("file already exists",
			zap.String("FileName", req.GetFile().Name),
			zap.Int64("FileSize", req.GetFile().Size),
			zap.String("Md5Sum", req.GetFile().Checksum))
		return nil
	}
	req.File.IdCode = utils.GenerateRandomCode()
	s.log.Info("receiving new file",
		zap.String("FileName", req.GetFile().Name),
		zap.Int64("FileSize", req.GetFile().Size),
		zap.String("Md5Sum", req.GetFile().Checksum),
		zap.String("IdCode", req.GetFile().IdCode),
	)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	readyResponse := &pb.UploadFileResponse{
		IdCode:       req.File.IdCode,
		ResponseCode: pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_READY,
	}
	if e := s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, readyResponse); e != nil {
		return e
	}

	if _, e := io.CopyN(file, conn, req.File.Size); e != nil {
		return fmt.Errorf("io copy error:%v", e)
	}
	resFileInfo, err := utils.FileInfo(filePath)
	if err != nil {
		return fmt.Errorf("get file info error:%v", err)
	}
	resultResp := &pb.UploadFileResponse{
		IdCode:       req.File.IdCode,
		ResponseCode: pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_SUCCESS,
	}
	if resFileInfo.Checksum != req.File.Checksum {
		s.log.Warn("wrong checksum",
			zap.String("fileName", req.File.Name),
			zap.String("clientMD5", req.File.Checksum),
			zap.String("ServerMD5", resFileInfo.Checksum))
		if e := os.Remove(filePath); e != nil {
			return e
		}
		resultResp.ResponseCode = pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_FAILED
	}
	s.log.Info("file received successfully",
		zap.String("Name", req.File.Name),
		zap.String("IDCode", req.File.IdCode))
	return s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, resultResp)
}
