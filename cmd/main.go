package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/moneta-sofia/API-GO.git/internal/domain"
	"github.com/moneta-sofia/API-GO.git/internal/user"
)

func main() {
	server := http.NewServeMux()

	db := user.DB{
		Users: []domain.User{
			{ID: 1, FirstName: "John", LastName: "Doe", Email: "doe@mail.com"},
			{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "smith@mail.com"},
			{ID: 3, FirstName: "Alice", LastName: "Johnson", Email: "johnson@mail.com"},
		},
		MaxUserID: 3,
	}

	loger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	repo := user.NewRepository(db, loger)
	service := user.NewService(loger, repo)
	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
