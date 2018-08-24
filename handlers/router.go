package handlers

import (
	"github.com/Abdujabbor/log-converter/repository"
	"github.com/julienschmidt/httprouter"
)

//Router extended httprouter.Router
type Router struct {
	httprouter.Router
	provider repository.Provider
}

//InitRoutes initializing router
func InitRoutes(provider repository.Provider) *Router {
	router := Router{
		provider: provider,
	}
	router.GET("/logs", list)
	router.GET("/logs/:page", list)
	return &router
}
