package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
)

func main() {
	fmt.Println("------------ image creation ")
	width, height := 100, 100
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := width; i >= 0; i-- {
		canvas.Set(i, height/2, color.RGBA{255, 0, 0, 255})
	}

	// fmt.Println(m.Bounds())
	// fmt.Println(m.At(0, 0).RGBA())
	tempFile, err := ioutil.TempFile("", "jpg")
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()
	png.Encode(tempFile, canvas)
	fmt.Println(tempFile.Name())

}
