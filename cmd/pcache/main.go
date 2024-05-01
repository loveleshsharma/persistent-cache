package main

import (
	"log"
	"net/http"
)

func main() {
	err := initObjects()
	if err != nil {
		log.Fatalln("error occurred", err)
		return
	}

	http.HandleFunc("/set", apiHandler.Set)
	http.HandleFunc("/get", apiHandler.Get)

	http.ListenAndServe(":8080", nil)
}
