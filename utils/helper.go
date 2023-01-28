package utils

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/Gictorbit/gofiler/proto/pb"
	"io"
	"os"
	"path/filepath"
)

const (
	RandomCodeLength = 10
)

func FileInfo(fPath string) (*pb.File, error) {
	fileExtension := filepath.Ext(fPath)
	fileStat, err := os.Stat(fPath)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s MD5 checksum is %x \n", file.Name(), hash.Sum(nil))
	return &pb.File{
		Name:      file.Name(),
		Size:      fileStat.Size(),
		Extension: fileExtension,
		Checksum:  fmt.Sprintf("%x", hash.Sum(nil)),
	}, nil
}

func GenerateRandomCode() string {
	b := make([]byte, RandomCodeLength)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}
