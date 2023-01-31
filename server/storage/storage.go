package storage

import (
	"errors"
	"fmt"
	"github.com/Gictorbit/gofiler/proto/pb"
	"github.com/Gictorbit/gofiler/utils"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Storage interface {
	SaveFile(file *pb.File, conn net.Conn) error
	FileInfo(code string) (*pb.File, error)
	GetFile(conn net.Conn, fileInfo *pb.File) error
}

type FileStorage struct {
	randomCodeLength int
	FileBoxPath      string
}

var (
	ErrWrongChecksum = errors.New("wrong md5 checksum")
	ErrFileNotFound  = errors.New("file not found")
	ErrInvalidName   = errors.New("file name is invalid")
)

func NewStorage(boxPath string, codeLength int) Storage {
	return &FileStorage{
		FileBoxPath:      boxPath,
		randomCodeLength: codeLength,
	}
}

// SaveFile reads file from connection and saves file to os fs
func (fs *FileStorage) SaveFile(file *pb.File, conn net.Conn) error {
	newFileName := fmt.Sprintf("%s_%s", file.IdCode, file.Name)
	fPath := path.Join(fs.FileBoxPath, strings.TrimSpace(newFileName))
	f, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer f.Close()
	receivedBytes, e := io.CopyN(f, conn, file.Size)
	if e != nil && !errors.Is(e, io.EOF) {
		return fmt.Errorf("io copy error:%v", e)
	}
	if receivedBytes != file.Size {
		return fmt.Errorf("recived byte size not equal to file size")
	}
	fileInfo, err := utils.FileInfo(fPath)
	if err != nil {
		return fmt.Errorf("get file info error:%v", err)
	}
	if fileInfo.Checksum != file.Checksum {
		if e := f.Close(); e != nil {
			return e
		}
		if e := os.Remove(fPath); e != nil {
			return e
		}
		return ErrWrongChecksum
	}
	return nil
}

// GetFile loads file using share code and writes file to tcp conn
func (fs *FileStorage) GetFile(conn net.Conn, fileInfo *pb.File) error {
	realName := fileInfo.GetIdCode() + "_" + fileInfo.Name
	fPath := filepath.Join(fs.FileBoxPath, realName)
	openedFile, err := os.Open(strings.TrimSpace(fPath)) // For read access.
	if err != nil {
		return err
	}
	defer openedFile.Close()
	sentBytes, e := io.CopyN(conn, openedFile, fileInfo.Size)
	if e != nil && !errors.Is(e, io.EOF) {
		return e
	}
	if sentBytes != fileInfo.Size {
		return fmt.Errorf("sent bytes not equal to file size")
	}
	return nil
}

// FileInfo returns
func (fs *FileStorage) FileInfo(code string) (*pb.File, error) {
	var realFileName string
	e := filepath.Walk(fs.FileBoxPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && strings.Contains(info.Name(), code) {
			realFileName = info.Name()
		}
		return nil
	})
	if e != nil {
		return nil, e
	}
	if realFileName == "" {
		return nil, ErrFileNotFound
	}
	fPath := filepath.Join(fs.FileBoxPath, realFileName)
	fileInfo, err := utils.FileInfo(fPath)
	if err != nil {
		return nil, err
	}
	splitName := strings.Split(realFileName, "_")
	if len(splitName) < 2 || splitName[0] != code {
		return nil, ErrInvalidName
	}
	fileInfo.Name = splitName[1]
	fileInfo.IdCode = splitName[0]
	return fileInfo, nil
}
