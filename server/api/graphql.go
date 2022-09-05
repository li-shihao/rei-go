package api

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"rei.io/rei/graph"
	"rei.io/rei/internal/database"
)

func GraphQLHandler(connStr string) *handler.Server {
	// Initialise db
	db := new(database.EntClient)
	db.Init("postgres", connStr)

	h := handler.New(graph.NewSchema(db.GetClient()))

	h.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(lru.New(1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return h
}

func PlaygroundQLHandler(endpoint string) http.HandlerFunc {
	//endpoint argument must be same as graphql handler path
	playgroundHandler := playground.Handler("GraphQL", endpoint)

	return playgroundHandler
}
