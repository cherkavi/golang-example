package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, _ := client.NewEnvClient()

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("image: %v   status: %v   ports:%v   labels: %v   \n\n", container.Image, container.Status, container.Ports, container.Labels)
	}
}
