package main

import (
	"context"
	"log"

	"github.com/vier21/pc-01-network-be/pkg/device"
)

func main() {
	dev, err := device.GetDevice(device.CiscoDev)
	if err != nil {
		log.Fatal(err)
		return
	}

	router := dev.InitRouter()
	router.AddACL(context.Background())
}
