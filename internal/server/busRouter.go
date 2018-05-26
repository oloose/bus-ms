package server

import (
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/file"
)

func NewBusRouter(mBusplanUrl string, mServer *Server) {
	busSubRouter := mServer.NewSubRouter("/bus")

	busSubRouter.Use(
		content.TypeNegotiator(content.JSON),
	)

	//serve bus plan
	busSubRouter.Get("/plan", file.Content(mBusplanUrl))
}
