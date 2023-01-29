package main

import (
	"fmt"
	"github.com/Gictorbit/gofiler/client/tcpclient"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"os"
)

var (
	HostAddress string
	PortNumber  uint
	FilePath    string
)

func main() {
	app := &cli.App{
		Name:  "client",
		Usage: "go file transfer client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Usage:       "host address",
				Value:       "127.0.0.1",
				DefaultText: "127.0.0.1",
				Destination: &HostAddress,
			},
			&cli.UintFlag{
				Name:        "port",
				Usage:       "server port number",
				Value:       2023,
				DefaultText: "2023",
				Aliases:     []string{"p"},
				Destination: &PortNumber,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "upload",
				Usage: "uploads a file to server",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:        "file",
						Usage:       "file path for upload",
						Required:    true,
						Aliases:     []string{"f"},
						Destination: &FilePath,
					},
				},
				Action: func(ctx *cli.Context) error {
					listenAddr := net.JoinHostPort(HostAddress, fmt.Sprintf("%d", PortNumber))
					log.Println("listen address is ", listenAddr)
					client := tcpclient.NewClient(listenAddr, log.Default())
					if e := client.Connect(); e != nil {
						log.Fatal(e)
					}
					if err := client.UploadFile(FilePath); err != nil {
						log.Println("failed to upload file", err)
						return err
					}
					client.Stop()
					return nil
				},
			},
		},
	}
	if e := app.Run(os.Args); e != nil {
		log.Println("failed to run app", e)
	}
}
