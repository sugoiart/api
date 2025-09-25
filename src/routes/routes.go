package routes

import (
	"art-api/src/handlers/art"
	"art-api/src/handlers/index"
	"github.com/julienschmidt/httprouter"
)

func InitRoutes(router *httprouter.Router) {
	router.GET("/", index.Index)
	router.GET("/v1", index.Index)
	router.GET("/v1/art", art.Index)
	router.GET("/v1/art/random", art.Random)
	router.GET("/v1/art/all", art.All)
	router.GET("/v1/art/sha/:sha", art.BySHA)
	router.GET("/v1/art/stats", art.Stats)
}
