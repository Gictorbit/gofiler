package tcpsrv

import (
	"bytes"
	"github.com/Gictorbit/gofiler/proto/pb"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func (s *Server) UploadFileHandler(req *pb.UploadFileRequest, conn net.Conn) error {

	file, err := os.Create(strings.TrimSpace(req.FileName))
	if err != nil {
		log.Fatal(err)
	}
	//TODO check exists
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
