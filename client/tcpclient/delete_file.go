package tcpclient

import (
	"fmt"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
)

// DeleteFile removes a file from server
func (c *Client) DeleteFile(shareCode string) error {
	request := &pb.DeleteFileRequest{
		IdCode: shareCode,
	}
	if e := c.SendRequest(pb.MessageType_MESSAGE_TYPE_DELETE_FILE_REQUEST, request); e != nil {
		return fmt.Errorf("sending delete file request failed: %v", e)
	}
	c.log.Println("sent delete file request...")

	deleteFileResp, err := utils.ReadMessageFromConn(c.conn, &pb.DeleteFileResponse{})
	if err != nil {
		return fmt.Errorf("delete file response failed: %v", err)
	}
	respCode := deleteFileResp.Message.ResponseCode
	if respCode != pb.ResponseCode_RESPONSE_CODE_OK {
		if respCode == pb.ResponseCode_RESPONSE_CODE_NOT_FOUND {
			c.log.Println("share code not found")
			c.PrintCode("NotFound")
			return nil
		}
		c.log.Println("server cant delete file")
		return nil
	}
	c.PrintFileInfo(deleteFileResp.Message.GetFile())
	resultResp, err := utils.ReadMessageFromConn(c.conn, &pb.DeleteFileResponse{})
	if err != nil {
		return fmt.Errorf("delete file response failed: %v", err)
	}
	if resultResp.Message.ResponseCode == pb.ResponseCode_RESPONSE_CODE_SUCCESS {
		c.log.Println("file removed successfully")
	}
	return nil
}
