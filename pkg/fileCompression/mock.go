package fileCompression

import "image"

type ImageCompressionMock struct {
}

var _ ImageCompressorInterface = &ImageCompressionMock{}

func (c *ImageCompressionMock) CreateSmall(buffer []byte) (image.Image, error) {
	return nil, nil
}

func (c *ImageCompressionMock) CreateMedium(buffer []byte) (image.Image, error) {
	return nil, nil
}

func (c *ImageCompressionMock) CreateLarge(buffer []byte) (image.Image, error) {
	return nil, nil
}
