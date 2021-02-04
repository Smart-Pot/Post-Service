package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"postservice/data"
	"postservice/endpoints"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandlers(e endpoints.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter().PathPrefix("/post/").Subrouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		e.GetMultiple,
		decodePostHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("GET").Path("/post/{id}").Handler(httptransport.NewServer(
		e.GetSingle,
		decodePostHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("DELETE").Path("/{id}").Handler(httptransport.NewServer(
		e.Delete,
		decodePostHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("POST").Path("/new").Handler(httptransport.NewServer(
		e.Create,
		decodeNewPostHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	return r
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodePostHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		// Handler error
	}

	return endpoints.PostRequest{
		ID: id,
	}, nil

}

func decodeNewPostHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var c data.Post

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		return nil, err
	}
	return endpoints.NewPostRequest{
		NewPost: c,
	}, nil

}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
