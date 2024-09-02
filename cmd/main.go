package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/moneta-sofia/API-GO.git/internal/user"
	"github.com/moneta-sofia/API-GO.git/pkg/bootstrap"
	"github.com/moneta-sofia/API-GO.git/pkg/handler"
)

func main() {
	server := http.NewServeMux()

	db := bootstrap.NewDB()
	loger := bootstrap.NewLogger()
	repo := user.NewRepository(db, loger)
	service := user.NewService(loger, repo)
	ctx := context.Background()

	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoints(ctx, service))

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
