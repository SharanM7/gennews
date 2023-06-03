package main

import (
	"hack_23/utils"
	"log"
	"sync"
)

func init() {
	utils.InitFlags()
}

func main() {
	server := utils.NewServer(utils.ViperConfig.Port)

	var wg sync.WaitGroup
	server.Start(&wg)

	log.Println("Server started")
	wg.Wait()
	log.Println("Server stopped")
	log.Println("End")
}
