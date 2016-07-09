package main

import (
	"github.com/styutnev/gotranslate/api"
	"log"
	"net/http"
)

func main() {
	router := api.CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}


