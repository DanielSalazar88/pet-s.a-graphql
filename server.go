package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"pet-s.a-graphql/graph"
	"pet-s.a-graphql/settings"
)

var s = settings.ReadYaml()

var defaultPort = fmt.Sprintf("%v", s.Handlers.Port)

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Configurar opciones CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permitir solicitudes desde cualquier origen
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	// Añadir CORS a los manejadores
	handlerWithCors := c.Handler(http.DefaultServeMux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, handlerWithCors))
}
