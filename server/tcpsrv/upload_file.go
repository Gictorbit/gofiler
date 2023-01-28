package tcpsrv

import (
	"bytes"
	"errors"
	"github.com/Gictorbit/gofiler/proto/pb"
	"io"
	"net"
	"os"
	"path"
	"strings"
)

func (s *Server) UploadFileHandler(req *pb.UploadFileRequest, conn net.Conn) (err error) {
	fileID := generateRandomCode()
	filePath := path.Join(s.storage, strings.TrimSpace(req.FileName))
	defer func() {
		resultResp := &pb.UploadFileResponse{
			FileId: fileID,
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
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	readyResponse := &pb.UploadFileResponse{
		FileId:       generateRandomCode(),
		ResponseCode: pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_READY,
	}
	if e := s.SendResponse(conn, pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE, readyResponse); e != nil {
		return e
	}
	var currentByte int64 = 0
	fileBuffer := make([]byte, 0)
	for err == nil || err != io.EOF {
		if _, e := conn.Read(fileBuffer); e != nil {
			return e
		}
		cleanedFileBuffer := bytes.Trim(fileBuffer, "\x00")
		_, err = file.WriteAt(cleanedFileBuffer, currentByte)
		if len(string(fileBuffer)) != len(string(cleanedFileBuffer)) {
			break
		}
		currentByte += PacketMaxByteLength
	}
	if e := file.Close(); e != nil {
		return e
	}
	return nil
}
