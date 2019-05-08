package httpserver

import (
	"log"
	"net/http"
	"strconv"
)

type route struct {
	url     string
	handler http.HandlerFunc
}

// Add new routes here
var routes = []route{
	{"/", http.HandlerFunc(fibonacciHandler)},
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
