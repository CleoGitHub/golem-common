package fileCompression

import "image"

type ImageCompressorInterface interface {
	CreateSmall(buffer []byte) (image.Image, error)
	CreateMedium(buffer []byte) (image.Image, error)
	CreateLarge(buffer []byte) (image.Image, error)
}
