package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"student-housing-backend/ent"
	"student-housing-backend/graph/generated"
	"student-housing-backend/graph"

	// _ "entgo.io/ent/cmd/ent"
	_ "github.com/99designs/gqlgen/cmd"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
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
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Client: client,}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
