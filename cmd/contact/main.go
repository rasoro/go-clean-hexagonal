package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rasoro/go-clean-hexagonal/adapter/postgres"
	"github.com/rasoro/go-clean-hexagonal/di"
	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()
	conn := postgres.OpenConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	contactService := di.ConfigContactDI(conn)

	router := mux.NewRouter()
	router.Handle("/contact", http.HandlerFunc(contactService.Create)).Methods("POST")
	router.Handle("/contact", http.HandlerFunc(contactService.Fetch)).Methods("GET")
	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
