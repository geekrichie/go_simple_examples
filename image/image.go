package main

import (
	"bytes"
	"fmt"
	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/chai2010/webp"
	"io/ioutil"
)

func main() {

	img, err := imgio.Open("image/input.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	//inverted := effect.Invert(img)
	width := img.Bounds().Max.X-img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	resized := transform.Resize(img, int(0.5*float64(width)), int(0.5*float64(height)), transform.Linear)
	fmt.Println(resized.Bounds())
	//rotated := transform.Rotate(resized, 45, nil)

	if err := imgio.Save("image/output.png", resized, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}

	//result := adjust.Brightness(img, 0.25)
	result := adjust.Hue(img,100)
	if err := imgio.Save("image/output.jpg", result, imgio.JPEGEncoder(100)); err != nil {
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer
	_ = webp.Encode(&buf, img, &webp.Options{Lossless: true})
	if err := ioutil.WriteFile("image/demo.webp",buf.Bytes(),0666 ); err !=nil {
		fmt.Println(err)
	}
	img.Bounds().Dx()
}
