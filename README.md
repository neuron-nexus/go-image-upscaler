# Neuron Nexus Image Upscale

Neuron Nexus Image Upscale is a Go framework for upscaling images. It allows you to frame the original image to fit a custom aspect ratio. The framework allows you to place the image in a frame with the original color (white) or in a frame of a custom color. Images can be imported in .png, .jpg and .gif. Export is available in .png and .jpg.

## Installation



```bash
go get -u go.mod
```

## Usage (v1.0.0)

```go
package main

import (
	"image/color"
	"log"
	"os"

	imageupscaler "github.com/neuron-nexus/go-image-upscaler"
)

func main() {
	upscaler := imageupscaler.New()
	err := upscaler.SetImage("./test.jpg")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("./test2.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//upscaler.Upscale(3, 2)
	upscaler.UpscaleWithColor(3, 2, color.RGBA{172, 25, 135, 255})
	upscaler.Render(imageupscaler.JPG, file, nil)
}



## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.