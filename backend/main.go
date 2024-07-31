package main

import (
	"log"
	"net/http"
	"os"
	"wolfort-games/inject"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := inject.InjectServer()

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	}).Handler(srv)
	http.Handle("/games-server/", playground.Handler("GraphQL playground", "/games-server/query"))
	http.Handle("/games-server/query", handler)
	log.Printf("connect to http://localhost:%s/games-server/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
