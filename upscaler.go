package imageupscaler

import (
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"os"
)

type Upscaler struct {
	img      image.Image
	template *image.RGBA
}

func New() *Upscaler {
	return &Upscaler{}
}

func (u *Upscaler) SetImage(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	u.img = img

	return nil
}

func (u *Upscaler) Upscale(widthRatio, heightRatio int) {
	clr := color.RGBA{255, 255, 255, 255}
	u.UpscaleWithColor(widthRatio, heightRatio, clr)
}

func (u *Upscaler) UpscaleWithColor(widthRatio, heightRatio int, clr color.RGBA) {
	newWidth, newHeight := u.calculateNewSize(3, 2)
	start, end := u.createEmptyImage(newWidth, newHeight, clr)
	u.fillTemplate(start, end)
}

func (u *Upscaler) Render(imgType ImageType, file io.Writer, o *jpeg.Options) error {
	switch imgType {
	case JPG:
		return jpeg.Encode(file, u.template, o)
	case PNG:
		return png.Encode(file, u.template)
	default:
		return ErrorIncorrectImageType
	}
}
