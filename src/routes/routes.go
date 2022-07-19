package routes

import (
	"github.com/julienschmidt/httprouter"
	"art-api/src/handlers/art"
	"art-api/src/handlers/index"
)

func InitRoutes(router *httprouter.Router) {
	router.GET("/", index.Index)
	router.GET("/v1", index.Index)
	router.GET("/v1/art", art.Index)
	router.GET("/v1/art/random", art.Random)
	router.GET("/v1/art/all", art.All)
}