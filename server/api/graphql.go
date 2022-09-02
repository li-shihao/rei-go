package api

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"rei.io/rei/graph"
	"rei.io/rei/internal/database"
)

func GraphQLHandler(connStr string) http.HandlerFunc {
	// Initialise db
	db := new(database.EntClient)
	db.Init("postgres", connStr)

	h := handler.NewDefaultServer(graph.NewSchema(db.GetClient()))
	return h.ServeHTTP
}

func PlaygroundQLHandler(endpoint string) http.HandlerFunc {
	//endpoint argument must be same as graphql handler path

	playgroundHandler := playground.Handler("GraphQL", endpoint)

	return playgroundHandler.ServeHTTP
}
