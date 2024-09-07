package fileStorage

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cleogithub/golem-common/pkg/merror"
)

type local struct {
	domain  string
	baseDir string
}

// Ensure at compilation Local implement FileStorage interface
var _ FileStorage = &local{}

func NewLocalStorage(domain string, baseDir string) (FileStorage, error) {
	fs := &local{
		domain:  strings.TrimSuffix(domain, "/"),
		baseDir: strings.TrimSuffix(strings.TrimPrefix(baseDir, "/"), "/"),
	}
	return fs, nil
}

// Save bytes as file with given name
func (fs *local) Save(dir string, filename string, data []byte) (string, error) {
	dir = strings.TrimSuffix(strings.TrimPrefix(dir, "/"), "/")
	fullDir := fmt.Sprintf("%s/%s", fs.baseDir, dir)

	if err := os.MkdirAll(fullDir, 0777); err != nil {
		return "", NewErrCreateDirectory(fullDir)
	}

	fullPath := fmt.Sprintf("./%s/%s", fullDir, filename)
	if _, err := os.Stat(fullPath); err == nil {
		return "", NewErrFileExist(fullPath)
	}

	//Get the file extension
	ext, err := GetExtension(data)
	if err != nil {
		return "", merror.Stack(err)
	}
	fullPath = fullPath + ext

	w, err := os.Create(fullPath)
	if err != nil {
		return "", merror.Stack(err)
	}
	defer w.Close()

	bufferReader := bytes.NewBuffer(data)
	_, err = io.Copy(w, bufferReader)
	if err != nil {
		return "", merror.Stack(err)
	}

	url := fmt.Sprintf("%s/%s/%s%s", fs.domain, fullDir, filename, ext)

	return url, nil
}

// Remove file name as filename
func (fs *local) Remove(url string) error {
	fullPath := fmt.Sprintf("./%s", strings.TrimPrefix(url, fs.domain))

	if _, err := os.Stat(fullPath); err != nil {
		return NewErrFileNotExist(fullPath)
	}

	if err := os.Remove(fullPath); err != nil {
		return merror.Stack(err)
	}

	return nil
}
