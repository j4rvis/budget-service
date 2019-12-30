package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/j4rvis/budget-service/config"
	"github.com/j4rvis/budget-service/transaction"
	_ "github.com/lib/pq"
)

func main() {
	config := config.New()

	router := chi.NewRouter()

	srv := http.Server{
		Handler: router,
		Addr:    config.ListenAddr,
	}

	transactionHandler := transaction.NewHandler()

	router.Get("/", transactionHandler.GetTransactions)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Service started at %s", config.ListenAddr)

	<-ctx.Done()

	srv.Close()
}
