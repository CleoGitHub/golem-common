package fileStorage

import (
	"errors"
	"strings"
)

var (
	ErrCreateDirectory    = errors.New("failed to create directory {directory}")
	ErrFileExist          = errors.New("file already exist in {filepath}")
	ErrFileNotExist       = errors.New("file does not exist in {filepath}")
	ErrUnknownStorageType = errors.New("StorageType unknown")
	ErrNoTypeDetected     = errors.New("no type detected")
)

// Return a error that wrap ErrCreateDirectory and add directory informations
func NewErrCreateDirectory(dir string) error {
	return errors.New(strings.ReplaceAll(ErrCreateDirectory.Error(), "{directory}", dir))
}

// Return a error that wrap ErrFileExist and add file informations
func NewErrFileExist(filepath string) error {
	return errors.New(strings.ReplaceAll(ErrCreateDirectory.Error(), "{filepath}", filepath))
}

// Return a error that wrap ErrFileNotExist and add file informations
func NewErrFileNotExist(filepath string) error {
	return errors.New(strings.ReplaceAll(ErrFileNotExist.Error(), "{filepath}", filepath))
}
