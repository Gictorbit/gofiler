package main

import (
	"fmt"
	"github.com/Gictorbit/gofiler/server/storage"
	"github.com/Gictorbit/gofiler/server/tcpsrv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	HostAddress string
	PortNumber  uint
	StoragePath string
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("create new logger failed:%v\n", err)
	}
	pwdPath, err := os.Getwd()
	if err != nil {
		logger.Error("get pwd failed", zap.Error(err))
	}
	defaultStoragePath := filepath.Join(pwdPath, "filebox")

	app := &cli.App{
		Name:  "server",
		Usage: "go file transfer server",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "starts server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "host",
						Usage:       "host address",
						Value:       "127.0.0.1",
						DefaultText: "127.0.0.1",
						Destination: &HostAddress,
						EnvVars:     []string{"HOST"},
					},
					&cli.UintFlag{
						Name:        "port",
						Usage:       "server port number",
						Value:       2023,
						DefaultText: "2023",
						Aliases:     []string{"p"},
						Destination: &PortNumber,
						EnvVars:     []string{"PORT"},
					},
					&cli.PathFlag{
						Name:        "storage",
						Usage:       "storage directory path to save files",
						Value:       defaultStoragePath,
						DefaultText: defaultStoragePath,
						Destination: &StoragePath,
						EnvVars:     []string{"STORAGE_PATH"},
					},
				},
				Action: func(ctx *cli.Context) error {
					listenAddr := net.JoinHostPort(HostAddress, fmt.Sprintf("%d", PortNumber))
					if _, e := os.Stat(StoragePath); os.IsNotExist(e) {
						if mkdirErr := os.MkdirAll(StoragePath, os.ModePerm); mkdirErr != nil {
							return mkdirErr
						}
					}
					logger.Info("storage directory is", zap.String("Path", StoragePath))
					fileStorage := storage.NewStorage(StoragePath, 10)
					server := tcpsrv.NewServer(listenAddr, logger, fileStorage)
					go server.Start()

					sigs := make(chan os.Signal, 1)
					signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
					<-sigs
					server.Stop()
					return nil
				},
			},
		},
	}

	if e := app.Run(os.Args); e != nil {
		logger.Error("failed to run app", zap.Error(e))
	}

}
