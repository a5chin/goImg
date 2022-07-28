package goimg

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

type ImgConf interface {
	Save()
	Flip() GoImg
	Gray()
}

type GoImg struct {
	path          string
	array         [][]color.Color
	height, width int
}

func LoadImage(path string) (img GoImg) {
	file, _ := os.Open(path)
	defer file.Close()

	src, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	img.path = path

	size := src.Bounds().Size()
	img.width, img.height = size.X, size.Y

	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, src.At(i, j))
		}
		img.array = append(img.array, y)
	}

	return
}

func (img *GoImg) Save(path string) {
	xlen, ylen := len(img.array), len(img.array[0])
	rect := image.Rect(0, 0, xlen, ylen)
	dst := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			dst.Set(x, y, img.array[x][y])
		}
	}

	file, err := os.Create(path)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()

	png.Encode(file, dst.SubImage(dst.Rect))
}

func (img *GoImg) Flip(direction string) {
	if direction == "h" {
		for i := 0; i < len(img.array)/2; i++ {
			z := len(img.array) - i - 1
			img.array[i], img.array[z] = img.array[z], img.array[i]
		}
	} else if direction == "v" {
		for _, col := range img.array {
			for i := 0; i < len(col)/2; i++ {
				z := len(col) - i - 1
				col[i], col[z] = col[z], col[i]
			}
		}
	}
}
