package imageupscaler

import (
	"image"
	"image/color"
	"sync"
)

func (u *Upscaler) getImageSize() (int, int) {
	return u.img.Bounds().Dx(), u.img.Bounds().Dy()
}

func (u *Upscaler) calculateNewSize(widthRatio, heightRatio int) (int, int) {

	width, height := u.getImageSize()

	var width2, height2 int

	if width*heightRatio > height*widthRatio {
		dw := (2*widthRatio - width%widthRatio) % (2 * widthRatio)
		width2 = width + dw
		height2 = width2 * heightRatio / widthRatio
	} else {
		dh := (2*heightRatio - height%heightRatio) % (2 * heightRatio)
		height2 = height + dh
		width2 = height2 * widthRatio / heightRatio
	}

	return width2, height2
}

func (u *Upscaler) createEmptyImage(width, height int, clr color.RGBA) (*image.Point, *image.Point) {

	origWidth, origHeight := u.getImageSize()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	difWidth := width - origWidth
	difHeight := height - origHeight

	difLeft := difWidth / 2
	difRigth := difLeft + difWidth%2

	difTop := difHeight / 2
	difBottom := difTop + difHeight%2

	wg := sync.WaitGroup{}

	for wid := 0; wid < width; wid++ {
		if wid < difLeft || wid >= width-difRigth {
			wg.Add(1)
			go func(wid int) {
				defer wg.Done()
				for hei := 0; hei <= height; hei++ {
					img.Set(wid, hei, clr)
				}
			}(wid)
		} else {
			wg.Add(2)
			go func(wid int) {
				defer wg.Done()
				for hei := 0; hei < difTop; hei++ {
					img.Set(wid, hei, clr)
				}
			}(wid)
			go func(wid int) {
				defer wg.Done()
				for hei := height - difBottom; hei < height; hei++ {
					img.Set(wid, hei, clr)
				}
			}(wid)
		}
	}

	wg.Wait()

	startPoint := &image.Point{X: difLeft, Y: difTop}
	endPoint := &image.Point{X: width - difRigth, Y: height - difBottom}

	u.template = img

	return startPoint, endPoint
}

func (u *Upscaler) fillTemplate(start, end *image.Point) {
	var (
		startX    = start.X
		startY    = start.Y
		endX      = end.X
		endY      = end.Y
		origImage = u.img
	)

	wg := sync.WaitGroup{}

	for x := startX; x < endX; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				u.template.Set(x, y, origImage.At(x-startX, y-startY))
			}
		}(x)
	}
	wg.Wait()
}
