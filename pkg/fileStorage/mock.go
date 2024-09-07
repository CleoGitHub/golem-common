package fileStorage

type mock struct {
}

// Ensure at compilation Local implement FileStorage interface
var _ FileStorage = &mock{}

func NewMockStorage() (FileStorage, error) {
	fs := &mock{}
	return fs, nil
}

// Save bytes as file with given name
func (fs *mock) Save(dir string, filename string, data []byte) (string, error) {
	return "", nil
}

// Remove file name as filename
func (fs *mock) Remove(url string) error {
	return nil
}
