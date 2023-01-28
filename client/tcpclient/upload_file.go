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
	fileBuffer := make([]byte, utils.PacketMaxByteLength)
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
		}
		return ErrServerNotReady
	}
	//file to read
	var currentByte int64 = 0
	file, err := os.Open(strings.TrimSpace(fPath)) // For read access.
	if err != nil {
		c.log.Println("failed to open file")
		return err
	}
	defer file.Close()
	c.log.Println("uploading file...")
	//read file until there is an error
	for err == nil || err != io.EOF {
		_, err = file.ReadAt(fileBuffer, currentByte)
		currentByte += utils.PacketMaxByteLength
		_, e := c.conn.Write(fileBuffer)
		if e != nil {
			c.log.Println("upload file chunk failed")
			return e
		}
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
