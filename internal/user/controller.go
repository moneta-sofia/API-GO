package user

import (
	"context"
	"errors"
	"fmt"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
		GetAll Controller
		Get    Controller
	}

	GetReq struct {
		ID uint64
	}
	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func MakeEndpoints(ctx context.Context, s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoints(s),
		GetAll: makeGetAllEndpoints(s),
		Get:    makeGetEndpoints(s),
	}
}

func makeCreateEndpoints(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateReq)
		fmt.Println(req)
		if req.FirstName == "" {
			return nil, errors.New("first name is required")
		}
		if req.LastName == "" {
			return nil, errors.New("last name is required")
		}
		if req.Email == "" {
			return nil, errors.New("email is required")
		}
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makeGetAllEndpoints(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}

func makeGetEndpoints(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetReq)
		user, err := s.Get(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		fmt.Println(req)
		return user, nil
	}
}
