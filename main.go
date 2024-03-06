package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", store)
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":3000", store)
	server.Run()

}
