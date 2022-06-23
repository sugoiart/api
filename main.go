package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"art-api/src/routes"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	router := httprouter.New()
	routes.InitRoutes(router)
	fmt.Println("Server is running on port: " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}