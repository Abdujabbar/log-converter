package handlers

import (
	"github.com/Abdujabbor/log-converter/repository"
	"github.com/julienschmidt/httprouter"
)

var provider repository.Provider

//InitRouter initilize method
func InitRouter(p repository.Provider) *httprouter.Router {
	router := httprouter.New()
	provider = p
	router.GET("/logs", list)
	router.GET("/logs/:page", list)
	return router
}
