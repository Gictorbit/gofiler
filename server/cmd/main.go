package main

import (
	"crypto/rand"
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"log"
	"os"
	"path/filepath"
)

var (
	HostAddress   string
	PortNumber    uint
	StoragePath   string
	AdminPassword string
)

const (
	RandomPasswordLength = 15
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
	app := &cli.App{
		Name:  "file-server",
		Usage: "go file transfer server",
		Commands: []*cli.Command{
			{
				Name: "start",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "host",
						Usage:       "host address",
						Value:       "localhost",
						Aliases:     []string{"h"},
						Destination: &HostAddress,
						EnvVars:     []string{"HOST"},
						Required:    true,
					},
					&cli.UintFlag{
						Name:        "port",
						Usage:       "server port number",
						Value:       2023,
						Aliases:     []string{"p"},
						Destination: &PortNumber,
						EnvVars:     []string{"PORT"},
						Required:    true,
					},
					&cli.PathFlag{
						Name:        "storage",
						Usage:       "storage directory path to save files",
						Value:       filepath.Join(pwdPath, "storage"),
						Destination: &StoragePath,
						EnvVars:     []string{"STORAGE_PATH"},
					},
					&cli.StringFlag{
						Name:        "password",
						Usage:       "set admin password",
						Value:       generateRandomPassword(),
						Aliases:     []string{"pass"},
						Destination: &AdminPassword,
						EnvVars:     []string{"ADMIN_PASSWORD"},
					},
				},
			},
		},

		Action: func(*cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func generateRandomPassword() string {
	b := make([]byte, RandomPasswordLength)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}
