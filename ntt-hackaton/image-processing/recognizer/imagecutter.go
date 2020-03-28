package recognizer

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"

	"../smartpark"
	"github.com/oliamb/cutter"
)

type SubImage struct {
	ID                  int32
	X, Y, Width, Height int
	PathToImage         string
	RecognizedValue     string
	Direction           smartpark.LicensePlateEvent_Direction
}

func (s SubImage) String() string {
	if len(s.RecognizedValue) > 0 {
		return fmt.Sprintf("id: %v  plate: %v direction: %v \n", s.ID, s.RecognizedValue, s.Direction)
	} else {
		return fmt.Sprintf("id: %v  \n", s.ID)
	}

}

func NewSubImage(ID int32, direction int32, X, Y, Width, Height int) *SubImage {
	returnValue := SubImage{
		Height:    Height,
		Width:     Width,
		X:         X,
		Y:         Y,
		ID:        ID,
		Direction: smartpark.LicensePlateEvent_Direction(direction),
	}
	return &returnValue
}

func createTempFile() (*os.File, error) {
	tempFile, err := ioutil.TempFile("", "jpg")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name())
	return os.OpenFile(tempFile.Name()+".jpg", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
}

func saveImageToTempFile(targetImage image.Image) (string, error) {
	file, err := createTempFile()
	if err != nil {
		return "", nil
	}
	defer file.Close()
	jpeg.Encode(file, targetImage, nil)
	return file.Name(), nil
}

func cropImage(mainImage image.Image, subImage SubImage) (image.Image, error) {
	if croppedImage, err := cutter.Crop(mainImage, cutter.Config{
		Width:  subImage.Width,
		Height: subImage.Height,
		Anchor: image.Point{subImage.X, subImage.Y},
		Mode:   cutter.TopLeft, // optional, default value
	}); err != nil {
		return nil, err
	} else {
		return croppedImage, nil
	}
}

func CutSubImages(file *os.File, subImages ...*SubImage) error {
	var mainImage, _, err = image.Decode(file)
	if err != nil {
		fmt.Printf("can't decode file: %v   %v \n", file.Name(), err)
		return err
	}

	for _, subImage := range subImages {
		var image, err = cropImage(mainImage, *subImage)
		if err != nil {
			return err
		}
		if subImage.PathToImage, err = saveImageToTempFile(image); err != nil {
			return err
		}
	}
	return nil
}
