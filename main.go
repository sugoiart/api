package main

import (
	"art-api/src/routes"
	"github.com/julienschmidt/httprouter"
	"github.com/syumai/workers"
	"net/http"
)

func main() {
	router := httprouter.New()
	routes.InitRoutes(router)

	workers.Serve(http.HandlerFunc(router.ServeHTTP))
}
