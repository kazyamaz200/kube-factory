package main

import (
	"log"

	"github.com/kyamazawa/glue-go/provider"
)

func main() {
	api := provider.NewAPI()
	ch := make(chan bool)
	go api.Activate()
	log.Println(<-ch)
}
