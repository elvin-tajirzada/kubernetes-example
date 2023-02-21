package main

import (
	"github.com/elvin-tacirzade/kubernetes-example/pkg/app"
	"log"
)

func main() {
	a, err := app.Init()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	a.Start()
	a.Shutdown()
}
