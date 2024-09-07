package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/moneta-sofia/API-GO.git/internal/user"
	"github.com/moneta-sofia/API-GO.git/pkg/transport"
	"github.com/moneta-sofia/response/response"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, endpoints user.Endpoints) {
	router.HandleFunc("/users/", UserServer(ctx, endpoints))
}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		log.Println(r.Method, ": ", url)
		path, pathSize := transport.Clean(url)

		params := make(map[string]string)
		if pathSize == 4 && path[2] != "" {
			params["userID"] = path[2]
		}

		params["token"] = r.Header.Get("Authorization")

		tran := transport.New(w, r, context.WithValue(ctx, "params", params))

		var end user.Controller
		var deco func(ctx context.Context, r *http.Request) (interface{}, error)

		switch r.Method {
		case http.MethodGet:
			switch pathSize {
			case 3:
				end = endpoints.GetAll
				deco = decodeGetAllUser
			case 4:
				end = endpoints.Get
				deco = decodeGetUser
			}
		case http.MethodPost:
			switch pathSize {
			case 3:
				end = endpoints.Create
				deco = decodeCreateUser
			}
		case http.MethodPatch:
			switch pathSize {
			case 4:
				end = endpoints.Update
				deco = decodeUpdateUser
			}
		}

		if end != nil && deco != nil {
			tran.Server(
				transport.Endpoint(end),
				deco,
				encodeResponse,
				encodeError,
			)
		} else {
			InvalidMethod(w)

		}
	}
}

func decodeGetAllUser(ctx context.Context, r *http.Request) (interface{}, error) {
	// params := ctx.Value("params").(map[string]string)
	// if err := tokenVerify(params["token"]); err != nil {
	// 	return nil, response.Unauthorized(err.Error())
	// }
	return nil, nil
}
func decodeGetUser(ctx context.Context, r *http.Request) (interface{}, error) {

	params := ctx.Value("params").(map[string]string)
	// if err := tokenVerify(params["token"]); err != nil {
	// 	return nil, response.Unauthorized(err.Error())
	// }
	id, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}
	return user.GetReq{
		ID: id,
	}, nil
}
func decodeUpdateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req user.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: %v", err.Error()))
	}

	params := ctx.Value("params").(map[string]string)
	if err := tokenVerify(params["token"]); err != nil {
		return nil, response.Unauthorized(err.Error())
	}

	id, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		log.Printf("Error parsing userID: %v", err)
		return nil, response.BadRequest(err.Error())
	}
	req.ID = id
	log.Printf("Decoded UpdateRequest: %+v", req) // Imprime el request completo
	return req, nil
}

func decodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	params := ctx.Value("params").(map[string]string)
	if err := tokenVerify(params["token"]); err != nil {
		return nil, response.Unauthorized(err.Error())
	}
	var req user.CreateReq
	fmt.Print(req)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request: %v", err))
	}
	return req, nil
}

func tokenVerify(token string) error {
	if os.Getenv("TOKEN") != token {
		return errors.New("Invalid token")
	}
	return nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	r := res.(response.Response)
	w.WriteHeader(r.StatusCode())

	return json.NewEncoder(w).Encode(res)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := err.(response.Response)

	w.WriteHeader(resp.StatusCode())
	_ = json.NewEncoder(w).Encode(resp)
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "Method doesn't exist"}`, status)
}
