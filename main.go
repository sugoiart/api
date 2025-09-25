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

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		router.ServeHTTP(w, r)
	})

	workers.Serve(handler)
}
