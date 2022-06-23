package main

import (
	"art-api/src/routes"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := httprouter.New()
	routes.InitRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}