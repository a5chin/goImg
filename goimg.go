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
	image         image.Image
	path          string
	height, width int
}

func (img *GoImg) String() string {
	s := fmt.Sprintf(
		"path: %s, size: (%d, %d)", img.path, img.width, img.height,
	)

	return s
}

func LoadImage(path string) GoImg {
	path, _ = filepath.Abs(path)
	file, _ := os.Open(path)
	defer file.Close()

	src, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	size := src.Bounds().Size()
	width, height := size.X, size.Y

	img := GoImg{
		image:  src,
		path:   path,
		height: height,
		width:  width,
	}

	return img
}

func (img *GoImg) Save(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()

	png.Encode(file, img.image)
}

func (img *GoImg) Flip(direction string) GoImg {
	canvas := image.NewRGBA(image.Rect(0, 0, img.width, img.height))

	if direction == "h" {
		// holizontal flip
		for y := 0; y < img.height; y++ {
			for x1 := 0; x1 < img.width/2; x1++ {
				x2 := img.width - x1 - 1
				canvas.Set(x1, y, img.image.At(x2, y))
				canvas.Set(x2, y, img.image.At(x1, y))
			}
		}
	} else if direction == "v" {
		// vertical flip
		for y1 := 0; y1 < img.height/2; y1++ {
			for x := 0; x < img.width; x++ {
				y2 := img.height - y1 - 1
				canvas.Set(x, y1, img.image.At(x, y2))
				canvas.Set(x, y2, img.image.At(x, y1))
			}
		}
	}

	dst := GoImg{
		image:  canvas,
		path:   img.path,
		height: img.height,
		width:  img.width,
	}

	return dst
}

func (img *GoImg) Gray() GoImg {
	canvas := image.NewGray(image.Rect(0, 0, img.width, img.height))

	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			canvas.Set(x, y, img.image.At(x, y))
		}
	}

	gray := GoImg{
		image:  canvas,
		path:   img.path,
		height: img.height,
		width:  img.width,
	}

	return gray
}
