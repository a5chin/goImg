package goimg_test

import (
	"goimg"
	"testing"
)

var (
	empty goimg.GoImg = goimg.GoImg{}
	path  string      = "assets/images/a5chin.png"
)

func TestLoad(t *testing.T) {
	img := goimg.LoadImage(path)

	if img == empty {
		t.Errorf("Cannot load %v.", path)
	}
}

func TestFlip(t *testing.T) {
	img := goimg.LoadImage(path)

	holizontal := img.Flip("h")
	vertical := img.Flip("v")
	if (holizontal == empty) || (vertical == empty) {
		t.Errorf("Cannot flip %v.", path)
	}
}

func TestGray(t *testing.T) {
	img := goimg.LoadImage(path)

	gray := img.Gray()
	if gray == empty {
		t.Errorf("%v cannot convert to gray.", path)
	}
}
