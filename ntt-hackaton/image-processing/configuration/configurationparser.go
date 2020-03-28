package configuration

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

// ConfigEntry element of the configuration
type ConfigEntry struct {
	X         int   `xml:"x,attr"`
	Y         int   `xml:"y,attr"`
	Width     int   `xml:"width,attr"`
	Height    int   `xml:"height,attr"`
	ID        int32 `xml:"id,attr"`
	Direction int32 `xml:"direction,attr"`
}

func (c ConfigEntry) String() string {
	return fmt.Sprintf("id: %v,  x: %v,   y: %v,   w: %v,  h: %v, direction: %v", c.ID, c.X, c.Y, c.Width, c.Height, c.Direction)
}

// Configuration aggregator
type Configuration struct {
	Entries []ConfigEntry `xml:"entry"`
}

func ReadConfiguration(path string) []ConfigEntry {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("can't read configuration from file: %v", path)
	}

	var configuration Configuration
	xml.Unmarshal(text, &configuration)
	return configuration.Entries
}
