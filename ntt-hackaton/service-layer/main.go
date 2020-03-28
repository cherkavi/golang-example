package main

import (
	"service-layer/imageprocessingserver"
	"service-layer/uiserver"
	"service-layer/datastore"
)

func main() {
	store := datastore.New()

	go uiserver.ListenUI(store)
	imageprocessingserver.ListenImageProcessing(store)
}