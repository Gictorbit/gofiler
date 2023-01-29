package tcpclient

import (
	"errors"
	"fmt"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
	"io"
	"os"
	"strings"
)

var (
	ErrServerNotReady = errors.New("server is not ready to receive file")
)

func (c *Client) UploadFile(fPath string) error {
	reqFile, err := utils.FileInfo(fPath)
	if err != nil {
		return err
	}
	c.PrintFileInfo(reqFile)
	uploadFileRequest := &pb.UploadFileRequest{File: reqFile}
	if e := c.SendRequest(pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_REQUEST, uploadFileRequest); e != nil {
		return e
	}

	readyMsg, err := utils.ReadMessageFromConn(c.conn, &pb.UploadFileResponse{})
	if err != nil {
		return err
	}
	if readyMsg.Message.ResponseCode != pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_READY {
		if readyMsg.Message.ResponseCode == pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_EXISTS {
			c.log.Println("file is already exists on server")
			return nil
		}
		return ErrServerNotReady
	}
	//file to read

	file, err := os.Open(strings.TrimSpace(fPath)) // For read access.
	if err != nil {
		c.log.Println("failed to open file")
		return err
	}
	defer file.Close()
	c.log.Println("uploading file...")
	if _, e := io.Copy(c.conn, file); e != nil {
		return e
	}
	return nil
}

func (c *Client) PrintFileInfo(file *pb.File) {
	fmt.Println("####### FileInfo #########")
	fmt.Printf("Name: %s\n", file.Name)
	fmt.Printf("Extension: %s\n", file.Extension)
	fmt.Printf("Size: %d\n", file.Size)
	fmt.Printf("CheckSum: %s\n", file.Checksum)
	fmt.Println("#########################")
}
