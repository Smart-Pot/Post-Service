package endpoints

import (
	"postservice/data"
	"postservice/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetSingle   endpoint.Endpoint
	GetMultiple endpoint.Endpoint
	Create      endpoint.Endpoint
	Delete      endpoint.Endpoint
}

type PostResponse struct {
	Posts   []*data.Post
	Success int32
	Message string
}

type PostRequest struct {
	ID string
}

type NewPostRequest struct {
	NewPost data.Post
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetSingle:   makeGetSingleEndpoint(s),
		GetMultiple: makeGetMultipleEndpoint(s),
		Create:      makeCreateEndpoint(s),
		Delete:      makeDeleteEndpoint(s),
	}
}
