package behavior

import (
	"fmt"

	"../../concurrency/channel"
)

func ChaineOfResponsibilities() {
	fmt.Println("---- channel pipeline, GOG chain of responsibilities  ")
	channel.ExampleChannelPipeline()
}
