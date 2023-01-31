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

// UploadFile uploads a file to server and prints share code
func (c *Client) UploadFile(fPath string) error {
	//file to read
	openedFile, err := os.Open(strings.TrimSpace(fPath)) // For read access.
	if err != nil {
		c.log.Println("failed to open file")
		return err
	}
	defer openedFile.Close()
	reqFile, err := utils.FileInfo(fPath)
	if err != nil {
		return err
	}
	c.PrintFileInfo(reqFile)

	if e := c.SendRequest(pb.MessageType_MESSAGE_TYPE_UPLOAD_FILE_REQUEST, &pb.UploadFileRequest{File: reqFile}); e != nil {
		return e
	}

	readyMsg, err := utils.ReadMessageFromConn(c.conn, &pb.UploadFileResponse{})
	if err != nil {
		return err
	}
	if readyMsg.Message.ResponseCode != pb.ResponseCode_RESPONSE_CODE_READY {
		return ErrServerNotReady
	}

	c.log.Println("uploading file...")
	sentBytes, e := io.CopyN(c.conn, openedFile, reqFile.Size)
	if e != nil && !errors.Is(e, io.EOF) {
		c.log.Println("io copy error:", e)
		return e
	}
	c.log.Println("sent bytes:", sentBytes)
	if sentBytes != reqFile.Size {
		return fmt.Errorf("sent bytes not equal to file size")
	}
	c.log.Println("file sent....")
	resultMsg, err := utils.ReadMessageFromConn(c.conn, &pb.UploadFileResponse{})
	if err != nil {
		return err
	}
	if resultMsg.Message.ResponseCode == pb.ResponseCode_RESPONSE_CODE_SUCCESS {
		c.log.Println("file uploaded successfully")
	} else {
		c.log.Println("upload file failed")
	}
	c.PrintCode(resultMsg.Message.GetShareCode())
	return nil
}

func (c *Client) PrintFileInfo(file *pb.File) {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
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

func (c *Client) PrintCode(code string) {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"file share code"})
	t.AppendRows([]table.Row{
		{code},
	})
	t.AppendSeparator()
	t.Render()
}
