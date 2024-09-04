package user

import (
	"context"
	"fmt"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
		GetAll Controller
		Get    Controller
		Update Controller
	}

	GetReq struct {
		ID uint64
	}
	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	UpdateRequest struct {
		ID        uint64
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
	}
)

func MakeEndpoints(ctx context.Context, s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Get:    makeGetEndpoint(s),
		Update: makeUpdateEndpoints(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateReq)
		fmt.Println(req)
		if req.FirstName == "" {
			return nil, ErrFirstNameRequired
		}
		if req.LastName == "" {
			return nil, ErrLastNameRequired
		}
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makeGetAllEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}

func makeGetEndpoint(s Service) Controller {
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
func makeUpdateEndpoints(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)

		if req.FirstName != nil && *req.FirstName == "" {
			return nil, ErrFirstNameRequired
		}
		if req.LastName == nil && *req.LastName == "" {
			return nil, ErrLastNameRequired
		}

		if err := s.Update(ctx, req.ID, req.FirstName, req.LastName, req.Email); err != nil {
			return nil, err
		}
		return nil, nil
	}
}
