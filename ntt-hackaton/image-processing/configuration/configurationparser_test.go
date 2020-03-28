package configuration

import (
	"fmt"
	"testing"
)

func TestImageRecognition(t *testing.T) {
	configEntries := ReadConfiguration("config-test.xml")
	for _, eachEntry := range configEntries {
		fmt.Println(eachEntry)
	}
	if len(configEntries) != 3 {
		t.Fatalf("can't read configuration properly")
	}
}
