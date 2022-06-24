package routes

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"art-api/src/handlers/art"
	"art-api/src/handlers/index"
)

func InitRoutes(router *httprouter.Router) {
	router.GET("/api/art", art.Index)
	router.GET("/api/art/random", art.Random)
	router.GET("/api/art/all", art.All)
}

func InitMainRoutes(mux *http.ServeMux, router *httprouter.Router) {
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.HandleFunc("/api", index.Index)
	mux.Handle("/api/", router)
}