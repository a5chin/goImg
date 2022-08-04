<div align="center">

<h1>GoImg</h1>

[![Gotest](https://github.com/a5chin/goImg/actions/workflows/gotest.yml/badge.svg)](https://github.com/a5chin/goImg/actions/workflows/gotest.yml) [![License](https://img.shields.io/pypi/l/ansicolortags.svg)](https://img.shields.io/pypi/l/ansicolortags.svg)

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

</div>

## Usage
Currently
- Flip
- Grayscale 

are implemented.

### Flip (Holizontal, Vertical)
```go
package main

import "goimg"

func main() {
	img := LoadImage("assets/images/cat.jpg")

    // (h)olizontal or (v)ertical
	fliped := img.Flip("v")
	fliped.Save("assets/images/fliped.png")
}
```
<img alt="fliped.png" src="assets/images/fliped.png">

### GrayScale
```go
package main

import "goimg"

func main() {
	img := LoadImage("assets/images/cat.jpg")
	gray := img.Gray()
	gray.Save("assets/images/gray.png")
}
```

<img alt="gray.png" src="assets/images/gray.png">

