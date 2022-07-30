package main

import "goimg"

func main() {
	img := goimg.LoadImage("assets/images/a5chin.png")

	fliped := img.Flip("v")
	gray := img.Gray()

	fliped.Save("aasets/images/fliped.png")
	gray.Save("aasets/images/gray.png")
}
