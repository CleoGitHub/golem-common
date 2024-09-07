package fileStorage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/cleogithub/golem-common/pkg/merror"
)

type gcs struct {
	bucket string
	client *storage.Client
}

// Ensure at compilation Gcs implement FileStorage interface
var _ FileStorage = &gcs{}

func NewGCSStorage(bucket string) (FileStorage, error) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, merror.Stack(err)
	}

	fs := &gcs{
		bucket: bucket,
		client: client,
	}
	return fs, nil
}

// Save bytes as file with given name
func (fs *gcs) Save(dir string, filename string, data []byte) (string, error) {
	r := bytes.NewReader(data)
	w := fs.client.Bucket(fs.bucket).Object(dir + "/" + filename).NewWriter(context.Background())
	w.CacheControl = "public, max-age=86400"
	w.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}

	if _, err := io.Copy(w, r); err != nil {
		return "", merror.Stack(err)
	}
	if err := w.Close(); err != nil {
		return "", merror.Stack(err)
	}

	// storage.SignedURL()

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s/%s", fs.bucket, dir, filename), nil
	// return "https://storage.googleapis.com" gexist/images/itemColor/df715c49-eabf-4eed-980c-102e37d104d7-small, nil
}

// Remove file name as filename
func (fs *gcs) Remove(path string) error {
	path = strings.TrimPrefix(path, fmt.Sprintf("https://storage.googleapis.com/%s/", fs.bucket))
	err := fs.client.Bucket(fs.bucket).Object(path).Delete(context.Background())
	if err != nil {
		return merror.Stack(err)
	}
	return nil
}
