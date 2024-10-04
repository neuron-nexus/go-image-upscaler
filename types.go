package imageupscaler

var (
	JPG = ImageType{imgType: "jpeg"}
	PNG = ImageType{imgType: "png"}
)

type ImageType struct {
	imgType string
}
