package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

type ImgOperator interface {
	String() string
	Save()
	Flip() GoImg
	Gray() GoImg
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

	size := src.Bounds().Size()
	width, height := size.X, size.Y

	img = GoImg{
		array:  make([][]color.Color, height),
		path:   path,
		height: height,
		width:  width,
	}

	for y := 0; y < height; y++ {
		row := make([]color.Color, width)
		for x := 0; x < width; x++ {
			row[x] = src.At(x, y)
		}
		img.array[y] = row
	}

	return
}

func (img *GoImg) String() (s string) {
	s = fmt.Sprintf(
		`path: %s, size: (%d, %d)`, img.path, img.width, img.height,
	)

	return
}

func (img *GoImg) Save(path string) {
	height, width := img.height, img.width
	rect := image.Rect(0, 0, width, height)
	dst := image.NewNRGBA(rect)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dst.Set(x, y, img.array[y][x])
		}
	}

	file, err := os.Create(path)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()

	png.Encode(file, dst.SubImage(dst.Rect))
}

func (img *GoImg) Flip(direction string) (dst GoImg) {
	height, width := img.height, img.width
	dst = GoImg{
		array:  make([][]color.Color, height),
		path:   img.path,
		height: height,
		width:  width,
	}

	if direction == "h" {
		for y := 0; y < height; y++ {
			col := make([]color.Color, width)
			for x := 0; x < width/2; x++ {
				z := width - x - 1
				col[x], col[z] = img.array[y][z], img.array[y][x]
			}
			dst.array[y] = col
		}
	} else if direction == "v" {
		for x1 := 0; x1 < height/2; x1++ {
			x2 := height - x1 - 1
			dst.array[x1], dst.array[x2] = img.array[x2], img.array[x1]
		}
	}

	return
}

func (img *GoImg) Gray() (gray GoImg) {
	height, width := img.height, img.width
	gray = GoImg{
		array:  make([][]color.Color, height),
		path:   img.path,
		height: height,
		width:  width,
	}

	return
}
