package texture

import (
	"image"
	"image/png"
	"os"
	"ray-tracing/utils"
	"ray-tracing/vec"
)

type ImageTexture struct {
	Image         image.Image
	Width, Height int
}

func NewImageTexture(fileName string) ImageTexture {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(fileName)

	if err != nil {
		println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	img, _, err2 := image.Decode(file)

	if err2 != nil {
		println("Error: Image could not be decoded")
		os.Exit(1)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	return ImageTexture{
		Width:  width,
		Height: height,
		Image:  img,
	}
}

func (it ImageTexture) Value(u float64, v float64, p vec.Vec3) vec.Color {
	itv := utils.NewInterval(0, 1)

	_u := itv.Clamp(u)
	_v := 1.0 - itv.Clamp(v)

	i := float64(it.Width) * _u
	j := float64(it.Height) * _v

	r, g, b, a := it.Image.At(int(i), int(j)).RGBA()
	fa := float64(a)

	return *vec.New(float64(r)/fa, float64(g)/fa, float64(b)/fa)
}
