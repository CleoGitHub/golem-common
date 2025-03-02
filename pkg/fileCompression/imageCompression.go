package fileCompression

import (
	"bytes"
	"image"

	"image/draw"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"

	"github.com/cleogithub/golem-common/pkg/merror"
	"github.com/disintegration/imaging"
)

type ImageCompressionConf struct {
	Small  int `yaml:"small"`
	Medium int `yaml:"medium"`
	Large  int `yaml:"large"`
}

type ImageCompressor struct {
	Conf ImageCompressionConf
}

var _ ImageCompressorInterface = &ImageCompressor{}

func (c *ImageCompressor) CreateSmall(buffer []byte) (image.Image, error) {
	return processImage(buffer, c.Conf.Small)
}

func (c *ImageCompressor) CreateMedium(buffer []byte) (image.Image, error) {
	return processImage(buffer, c.Conf.Medium)
}

func (c *ImageCompressor) CreateLarge(buffer []byte) (image.Image, error) {
	return processImage(buffer, c.Conf.Large)
}

func processImage(buffer []byte, width int) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewBuffer(buffer))
	if err != nil {
		return nil, merror.Stack(err)
	}

	var res *image.NRGBA
	if width == 0 {
		res = imaging.Clone(img)
	} else {
		res = imaging.Resize(img, width, 0, imaging.Lanczos)
	}

	if paletted, ok := img.(*image.Paletted); ok {
		res := image.NewPaletted(res.Bounds(), paletted.Palette)
		draw.Src.Draw(res, res.Bounds(), res, res.Bounds().Min)
	}

	return res, err
}
