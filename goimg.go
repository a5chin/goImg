package goimg

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

type ImgOperator interface {
	Gray() GoImg
	Flip() GoImg
	Save()
	String() string
}

type GoImg struct {
	Image         image.Image
	Path          string
	Height, Width int
}

func (img *GoImg) String() string {
	s := fmt.Sprintf(
		"path: %s, size: (%d, %d)", img.Path, img.Width, img.Height,
	)

	return s
}

func LoadImage(path string) (img GoImg) {
	path, _ = filepath.Abs(path)
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Println("Cannot open file:", err)
	}

	src, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	size := src.Bounds().Size()
	width, height := size.X, size.Y

	img = GoImg{
		Image:  src,
		Path:   path,
		Height: height,
		Width:  width,
	}

	return img
}

func (img *GoImg) Save(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()

	png.Encode(file, img.Image)
}

func (img *GoImg) Flip(direction string) GoImg {
	canvas := image.NewRGBA(
		image.Rect(0, 0, img.Width, img.Height),
	)

	if direction == "h" {
		// holizontal flip
		for y := 0; y < img.Height; y++ {
			for x1 := 0; x1 < img.Width/2; x1++ {
				x2 := img.Width - x1 - 1
				canvas.Set(x1, y, img.Image.At(x2, y))
				canvas.Set(x2, y, img.Image.At(x1, y))
			}
		}
	} else if direction == "v" {
		// vertical flip
		for y1 := 0; y1 < img.Height/2; y1++ {
			for x := 0; x < img.Width; x++ {
				y2 := img.Height - y1 - 1
				canvas.Set(x, y1, img.Image.At(x, y2))
				canvas.Set(x, y2, img.Image.At(x, y1))
			}
		}
	}

	dst := GoImg{
		Image:  canvas,
		Path:   img.Path,
		Height: img.Height,
		Width:  img.Width,
	}

	return dst
}

func (img *GoImg) Gray() GoImg {
	canvas := image.NewGray(
		image.Rect(0, 0, img.Width, img.Height),
	)

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			canvas.Set(x, y, img.Image.At(x, y))
		}
	}

	gray := GoImg{
		Image:  canvas,
		Path:   img.Path,
		Height: img.Height,
		Width:  img.Width,
	}

	return gray
}
