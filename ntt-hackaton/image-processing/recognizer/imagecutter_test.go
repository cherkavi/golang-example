package recognizer

import (
	"fmt"
	"os"
	"testing"
)

func TestCutter(t *testing.T) {
	var file, err = os.OpenFile("test-image.jpg", os.O_RDONLY, 400)
	if err != nil {
		t.Error(err)
	}
	var subImages []*SubImage
	subImages = append(subImages, NewSubImage(10, 10, 100, 100))
	subImages = append(subImages, NewSubImage(100, 100, 100, 100))
	subImages = append(subImages, NewSubImage(200, 200, 100, 100))
	CutSubImages(file, subImages...)

	for _, image := range subImages {
		fmt.Println(image.PathToImage)
	}
}
