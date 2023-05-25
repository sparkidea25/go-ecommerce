package main

import (
	"net/http"

	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()



	server := &http.Server{
		Addr: "0.0.0.0.8008",

	}

	server.ListenAndServe()
}

