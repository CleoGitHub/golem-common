package fileStorage

import (
	"mime"
	"net/http"

	"github.com/cleogithub/golem-common/pkg/merror"
	"github.com/google/uuid"
)

type StorageType string

const (
	LOCAL StorageType = "local"
	GCS   StorageType = "gcs"
)

type FileStorage interface {
	//Save bytes as file with given name
	Save(dir string, filename string, data []byte) (string, error)

	//Remove file name as filename
	Remove(path string) error
}

func GenerateFilename() string {
	return uuid.NewString()
}

func GetExtension(data []byte) (string, error) {
	exts, err := mime.ExtensionsByType(http.DetectContentType(data))
	if err != nil {
		return "", merror.Stack(err)
	}

	if len(exts) == 0 {
		return "", merror.Stack(ErrNoTypeDetected)
	}

	return exts[0], nil
}
