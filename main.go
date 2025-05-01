package main

import (
	"log"
	"os"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal("Error while spawning a pg instance: ", err)
		os.Exit(0)
	}

	//	fmt.Printf("%+v\n", store) // %v+\n gives super verbose output :D
	server := NewApiServer(":3000", store)
	server.Run()
}
