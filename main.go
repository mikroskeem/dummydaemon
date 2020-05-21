package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	log.Println("running")
	<-sigCh
	log.Println("got signal, exiting")
}
