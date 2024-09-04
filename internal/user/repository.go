package user

import (
	"context"
	"errors"
	"log"
	"slices"

	"github.com/moneta-sofia/API-GO.git/internal/domain"
)

type DB struct {
	Users     []domain.User
	MaxUserID uint64
}

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
		Get(ctx context.Context, id uint64) (*domain.User, error)
	}
	repo struct {
		db  DB
		log *log.Logger
	}
)

func NewRepository(db DB, log *log.Logger) Repository {
	return &repo{
		db:  db,
		log: log,
	}
}

func (r *repo) Create(ctx context.Context, user *domain.User) error {
	r.db.MaxUserID++
	user.ID = r.db.MaxUserID
	r.db.Users = append(r.db.Users, *user)
	r.log.Println("Repository created")
	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	r.log.Println("Repository get all")
	return r.db.Users, nil
}

func (r *repo) Get(ctx context.Context, id uint64) (*domain.User, error) {
	index := slices.IndexFunc(r.db.Users, func(v domain.User) bool {
		return v.ID == id
	})
	if index < 0 {
		return nil, errors.New("users doesnt't exist")
	}
	return &r.db.Users[index], nil
}
