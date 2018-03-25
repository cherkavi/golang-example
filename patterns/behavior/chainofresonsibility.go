package behavior

import (
	"fmt"

	"../../concurrency/channel"
)

func ChaineOfResponsibilities() {
	fmt.Println("---- channel pipeline, GOF chain of responsibilities  ")
	channel.ExampleChannelPipeline()
}
