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
	"strings"
)

type Storage interface {
	SaveFile(file *pb.File, conn net.Conn) error
}

type FileStorage struct {
	randomCodeLength int
	FileBoxPath      string
}

var (
	ErrWrongChecksum = errors.New("wrong md5 checksum")
)

func NewStorage(boxPath string, codeLength int) Storage {
	return &FileStorage{
		FileBoxPath:      boxPath,
		randomCodeLength: codeLength,
	}
}

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
	fileInfo, err := utils.FileInfo(fPath, f)
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
