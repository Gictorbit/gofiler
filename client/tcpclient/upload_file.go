package tcpclient

import (
	"errors"
	"fmt"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
	"github.com/jedib0t/go-pretty/v6/table"
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
	io.CopyN(c.conn, file, reqFile.Size)

	c.log.Println("file sent....")
	resultMsg, err := utils.ReadMessageFromConn(c.conn, &pb.UploadFileResponse{})
	if err != nil {
		return err
	}
	if resultMsg.Message.ResponseCode == pb.UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_SUCCESS {
		c.log.Println("file uploaded successfully")
	} else {
		c.log.Println("upload file failed")
	}
	fmt.Println("File ID Code is: ", resultMsg.Message.GetIdCode())
	return nil
}

func (c *Client) PrintFileInfo(file *pb.File) {
	fmt.Println("FileInfo:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"file", "info"})
	t.AppendRows([]table.Row{
		{"Name", file.Name},
		{"Extension", file.Extension},
		{"Size", fmt.Sprintf("%d Bytes", file.Size)},
		{"CheckSum", file.Checksum},
	})
	t.AppendSeparator()
	t.Render()
}
