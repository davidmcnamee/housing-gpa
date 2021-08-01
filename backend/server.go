package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"student-housing-backend/ent"
	"student-housing-backend/graph/generated"
	"student-housing-backend/graph/schema"

	_ "github.com/99designs/gqlgen/cmd"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func getDbClient() *ent.Client {
	var client *ent.Client
	var err error
	if os.Getenv("ENV") == "production" {
		client, err = ent.Open("postgres", os.Getenv("DATABASE_URL"))
	} else {
		client, err = ent.Open("sqlite3", "file:db.sqlite3?_fk=1")
	}
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func main() {
	client := getDbClient()
	defer client.Close()

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:"+port},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &schema.Resolver{Client: client,}}))
	router.Handle("/", playground.Handler("Student Housing", "/query"))
	router.Handle("/query", srv)
	
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
