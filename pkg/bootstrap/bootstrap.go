package bootstrap

import (
	"log"
	"os"

	"github.com/moneta-sofia/API-GO.git/internal/domain"
	"github.com/moneta-sofia/API-GO.git/internal/user"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() user.DB {
	return user.DB{
		Users: []domain.User{
			{ID: 1, FirstName: "John", LastName: "Doe", Email: "doe@mail.com"},
			{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "smith@mail.com"},
			{ID: 3, FirstName: "Alice", LastName: "Johnson", Email: "johnson@mail.com"},
		},
		MaxUserID: 3,
	}
}
