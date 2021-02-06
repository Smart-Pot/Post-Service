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
	Vote        endpoint.Endpoint
}

type PostResponse struct {
	Posts   []*data.Post
	Success int32
	Message string
}

type VoteRequest struct {
	UserID string `json:"userId"`
	PostID string `json:"postId"`
}

type PostRequest struct {
	ID     string
	UserID string
}

type NewPostRequest struct {
	NewPost data.Post
	UserID  string
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetSingle:   makeGetSingleEndpoint(s),
		GetMultiple: makeGetMultipleEndpoint(s),
		Create:      makeCreateEndpoint(s),
		Delete:      makeDeleteEndpoint(s),
		Vote:        makeVoteEndpoint(s),
	}
}
