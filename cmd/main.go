package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/moneta-sofia/API-GO.git/internal/user"
	"github.com/moneta-sofia/API-GO.git/pkg/bootstrap"
	"github.com/moneta-sofia/API-GO.git/pkg/handler"
)

func main() {
	_ = godotenv.Load()

	server := http.NewServeMux()

	db, err := bootstrap.NewDB()

	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}

	loger := bootstrap.NewLogger()
	repo := user.NewRepository(db, loger)
	service := user.NewService(loger, repo)
	ctx := context.Background()

	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoints(ctx, service))
	port := os.Getenv("PORT")

	fmt.Println("Server started at port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}
