package tcpclient

import (
	"fmt"
	"github.com/Gictorbit/gofiler/utils"
	"log"
	"net"
	"sync"
)

type Client struct {
	listenAddr string
	conn       net.Conn
	wg         sync.WaitGroup
	log        *log.Logger
}
type ClientInterface interface {
	UploadFile(fPath string) error
	DownloadFile(outPath, shareCode string) error
	DeleteFile(shareCode string) error
	FileInfo(shareCode string) error
}

var (
	_ ClientInterface = &Client{}
)

func NewClient(listenAddr string, logger *log.Logger) *Client {
	return &Client{
		listenAddr: listenAddr,
		wg:         sync.WaitGroup{},
		log:        logger,
	}
}

func (c *Client) Connect() error {
	conn, err := net.Dial(utils.ServerSocketType, c.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to dial server: %v\n", err.Error())
	}
	c.conn = conn
	return nil
}

func (c *Client) Stop() {
	c.wg.Wait()
	c.conn.Close()
	c.log.Println("stop client...")
}
