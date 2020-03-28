package recognizer

import (
	"fmt"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	var file, err = os.OpenFile("test-image-single.jpg", os.O_RDONLY, 400)
	output, err := RecognizeImage(file.Name())
	if err != nil {
		t.Error(err)
		t.Error(output)
	}
	fmt.Println(output)
}

func TestParserAndCutter(t *testing.T) {
	var file, err = os.OpenFile("test-image.jpg", os.O_RDONLY, 400)
	if err != nil {
		t.Error(err)
	}
	var subImages []*SubImage
	subImages = append(subImages, NewSubImage(0, 0, 300, 768))
	subImages = append(subImages, NewSubImage(300, 000, 300, 768))
	subImages = append(subImages, NewSubImage(700, 000, 330, 768))
	CutSubImages(file, subImages...)

	for _, image := range subImages {
		image.RecognizedValue, err = RecognizeImage(image.PathToImage)
		fmt.Printf("path: %v \n", image.PathToImage)
		fmt.Printf("value: %v \n", image.RecognizedValue)
		fmt.Printf("err: %v \n", err)
	}
}
