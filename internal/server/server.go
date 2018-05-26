package server

import (
	"log"
	"net/http"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/access"
	"github.com/go-ozzo/ozzo-routing/fault"
	"github.com/go-ozzo/ozzo-routing/file"
	"github.com/go-ozzo/ozzo-routing/slash"
)

const (
	port = ":8080"
)

// Defines a Server type containing a router that defines the rest-api routing.
type Server struct {
	router *routing.Router
}

func NewServer(mBusplanUrl string) *Server {
	server := Server{router: routing.New()}
	// define base router config
	server.router.Use(
		access.Logger(log.Printf),
		slash.Remover(http.StatusMovedPermanently),
		fault.Recovery(log.Printf),
	)

	// server swagger-ui
	server.router.Get("/swagger/*", file.Server(file.PathMap{
		"/swagger/": "/swagger-ui/dist/",
	}))

	// add sup routes (/bus/*)
	NewBusRouter(mBusplanUrl, &server)

	return &server
}

func (rServer *Server) NewSubRouter(mPath string) *routing.RouteGroup {
	return rServer.router.Group(mPath)
}

func (rServer *Server) Start() {
	log.Println("Listining on port " + port)
	// add api routes
	http.Handle("/", rServer.router)
	// start server
	http.ListenAndServe(port, nil)
}
