package tcpclient

import (
	"github.com/Gictorbit/gofiler/utils"
	"log"
	"net"
	"sync"
)

type Empty struct{}

type Client struct {
	listenAddr string
	conn       net.Conn
	quitChan   chan Empty
	wg         sync.WaitGroup
	log        *log.Logger
}

func NewClient(listenAddr string, logger *log.Logger) *Client {
	return &Client{
		listenAddr: listenAddr,
		quitChan:   make(chan Empty),
		wg:         sync.WaitGroup{},
		log:        logger,
	}
}

func (c *Client) Start() {
	conn, err := net.Dial(utils.ServerSocketType, c.listenAddr)
	if err != nil {
		c.log.Println("failed to dial server", err.Error())
		return
	}
	defer conn.Close()
	c.conn = conn
	<-c.quitChan
}

func (c *Client) Stop() {
	c.wg.Wait()
	c.quitChan <- Empty{}
	c.log.Println("stop client...")
}
