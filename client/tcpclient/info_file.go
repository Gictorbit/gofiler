package tcpclient

import (
	"fmt"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
)

// FileInfo fetches file info from server using share code
func (c *Client) FileInfo(shareCode string) error {
	request := &pb.FileInfoRequest{
		ShareCode: shareCode,
	}
	if e := c.SendRequest(pb.MessageType_MESSAGE_TYPE_FILE_INFO_REQUEST, request); e != nil {
		return fmt.Errorf("sending file info request failed: %v", e)
	}
	c.log.Println("sent file info request...")

	fileInfo, err := utils.ReadMessageFromConn(c.conn, &pb.FileInfoResponse{})
	if err != nil {
		return fmt.Errorf("get file info response failed: %v", err)
	}
	switch fileInfo.Message.ResponseCode {
	case pb.ResponseCode_RESPONSE_CODE_SUCCESS:
		c.PrintFileInfo(fileInfo.Message.GetFileInfo())
	case pb.ResponseCode_RESPONSE_CODE_NOT_FOUND:
		c.log.Println("share code not found")
		c.PrintCode("NotFound")
	case pb.ResponseCode_RESPONSE_CODE_FAILED:
		c.log.Println("get file info failed!!")
	}
	return nil
}
