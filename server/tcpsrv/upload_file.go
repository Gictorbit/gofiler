package tcpsrv

import (
	"errors"
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
	reqFile := req.File
	reqFile.IdCode = utils.GenerateRandomCode()
	s.log.Info("receiving new file", zap.Any("File", reqFile))
	filePath := path.Join(s.storage, strings.TrimSpace(reqFile.Name))
	defer func() {
		resultResp := &pb.UploadFileResponse{
			IdCode: reqFile.IdCode,
		}
		if err != nil {
			resultResp.ResponseCode = pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_FAILED
		} else {
			resultResp.ResponseCode = pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_SUCCESS
		}
		_ = s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, resultResp)
	}()

	if _, err = os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
		existsResp := &pb.UploadFileResponse{
			ResponseCode: pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_EXISTS,
		}
		if e := s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, existsResp); e != nil {
			return e
		}
		s.log.Warn("file already exists", zap.Any("File", req.GetFile()))
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	readyResponse := &pb.UploadFileResponse{
		IdCode:       reqFile.IdCode,
		ResponseCode: pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_READY,
	}
	if e := s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, readyResponse); e != nil {
		return e
	}
	if _, e := io.Copy(file, conn); e != nil && e != io.EOF {
		return e
	}
	return file.Close()
}
