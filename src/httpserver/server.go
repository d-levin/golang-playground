package httpserver

import (
	"fibonacci"
	"log"
	"net/http"
	"strconv"
	"todos"
)

type route struct {
	url     string
	handler http.HandlerFunc
}

// Add new routes here
var routes = []route{
	{"/", http.HandlerFunc(fibonacci.Handler)},
	{"/todos/", http.HandlerFunc(todos.Handler)},
}

func Start(port int) {
	addr := ":" + strconv.Itoa(port)

	registerRoutes()

	log.Fatalln(http.ListenAndServe(addr, nil))
}

func registerRoutes() {
	for _, route := range routes {
		http.Handle(route.url, route.handler)
	}
}
