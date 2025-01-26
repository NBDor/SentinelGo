package main

import (
	"log"
	"net/http"
)

type Plugin interface {
	Name() string
	Process(http.Handler) http.Handler
}

type Gateway struct {
	plugins []Plugin
	router  *http.ServeMux
}

func NewGateway() *Gateway {
	return &Gateway{
		plugins: make([]Plugin, 0),
		router:  http.NewServeMux(),
	}
}

func (g *Gateway) AddPlugin(p Plugin) {
	g.plugins = append(g.plugins, p)
}

func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler = g.router

	// Apply plugins in reverse order
	for i := len(g.plugins) - 1; i >= 0; i-- {
		handler = g.plugins[i].Process(handler)
	}

	handler.ServeHTTP(w, r)
}

func main() {
	gateway := NewGateway()

	// Example route
	gateway.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("SentinelGo API Gateway"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: gateway,
	}

	log.Printf("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
