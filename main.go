package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sparkidea25/go-postgres/models"

)


func main() {
	godotenv.Load()



	server := &http.Server{
		Addr: "0.0.0.0.8008",

	}

	models.connectDatabase();

	server.ListenAndServe()

}

