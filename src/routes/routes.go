package routes

import (
	"github.com/julienschmidt/httprouter"
	"art-api/src/handlers/art"
	"art-api/src/handlers/index"
)

func InitRoutes(router *httprouter.Router) {
	router.GET("/api", index.Index)
	router.GET("/api/random", art.Random)
	router.GET("/api/all", art.All)
}