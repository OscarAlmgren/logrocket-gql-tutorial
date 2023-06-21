package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/oscaralmgren/logrocket-gql-tutorial/graph"
	loggy "github.com/oscaralmgren/logrocket-gql-tutorial/log"
)

const defaultPort = "8080"

func init() {
	loggy.InitLog()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	loggy.Logger.Info().Msgf("connect to http://localhost:%s/ for GraphQL playground", port)
	httpserver := &http.Server{Addr: ":" + port}
	if err := httpserver.ListenAndServe(); err != nil {
		loggy.Logger.Err(err).Msg("Failed to start the http server")
	}
}
