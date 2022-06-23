package routes

import (
	"github.com/julienschmidt/httprouter"
	"art-api/src/handlers/art"
)

func InitRoutes(router *httprouter.Router) {
	router.GET("/api/random", art.Random)
	router.GET("/api/all", art.All)
}