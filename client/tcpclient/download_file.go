package tcpclient

import (
	"errors"
	"fmt"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
	"io"
	"os"
	"path"
	"strings"
)

// DownloadFile downloads a file from server by its share code
func (c *Client) DownloadFile(outPath, shareCode string) error {
	request := &pb.DownLoadFileRequest{
		ShareCode: shareCode,
	}
	if e := c.SendRequest(pb.MessageType_MESSAGE_TYPE_DOWNLOAD_FILE_REQUEST, request); e != nil {
		return fmt.Errorf("sending download request failed: %v", e)
	}
	c.log.Println("sent download file request...")

	downloadFileResp, err := utils.ReadMessageFromConn(c.conn, &pb.DownLoadFileResponse{})
	if err != nil {
		return fmt.Errorf("get file download response failed: %v", err)
	}
	respCode := downloadFileResp.Message.ResponseCode
	if respCode != pb.ResponseCode_RESPONSE_CODE_READY {
		if respCode == pb.ResponseCode_RESPONSE_CODE_NOT_FOUND {
			c.log.Println("share code not found")
			c.PrintCode("NotFound")
			return nil
		}
		c.log.Println("server is not ready to receive file")
		return nil
	}
	c.PrintFileInfo(downloadFileResp.Message.GetFile())
	fPath := path.Join(outPath, strings.TrimSpace(downloadFileResp.Message.GetFile().Name))
	f, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer f.Close()
	receivedBytes, e := io.CopyN(f, c.conn, downloadFileResp.Message.GetFile().GetSize())
	if e != nil && !errors.Is(e, io.EOF) {
		return fmt.Errorf("io copy error:%v", e)
	}
	if receivedBytes != downloadFileResp.Message.GetFile().GetSize() {
		return fmt.Errorf("recived byte size not equal to file size")
	}
	resultResp, err := utils.ReadMessageFromConn(c.conn, &pb.DownLoadFileResponse{})
	if err != nil {
		return fmt.Errorf("get file download response failed: %v", err)
	}
	if resultResp.Message.ResponseCode == pb.ResponseCode_RESPONSE_CODE_SUCCESS {
		c.log.Println("file downloaded successfully", fPath)
	}
	fileInfo, err := utils.FileInfo(fPath)
	if err != nil {
		return fmt.Errorf("get file info error:%v", err)
	}
	if fileInfo.Checksum != downloadFileResp.Message.GetFile().GetChecksum() {
		if e := f.Close(); e != nil {
			return e
		}
		if e := os.Remove(fPath); e != nil {
			return e
		}
		return fmt.Errorf("wrong md5 checksum")
	}
	return nil
}
