package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	err := http.ListenAndServe(":3030", middlewares.RemoveTrailingSlash(router))
	if err != nil {
		panic(err)
	}
}
