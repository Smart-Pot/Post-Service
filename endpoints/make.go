package endpoints

import (
	"context"
	"postservice/service"

	"github.com/go-kit/kit/endpoint"
)

func makeGetSingleEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		result, err := s.GetSingle(ctx, req.ID)
		response := PostResponse{Posts: result, Success: 1, Message: "Posts found!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}

func makeGetMultipleEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostsRequest)
		result, err := s.GetMultiple(ctx, req.ID, req.PageNumber, req.PageSize)
		response := PostResponse{Posts: result, Success: 1, Message: "Posts found!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}

func makeVoteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VoteRequest)
		err := s.Vote(ctx, req.UserID, req.PostID)
		response := PostResponse{Posts: nil, Success: 1, Message: "Vote successful!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NewPostRequest)
		err := s.Create(ctx, req.UserID, req.NewPost)
		response := PostResponse{Posts: nil, Success: 1, Message: "Post Created!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}
func makeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		err := s.Delete(ctx, req.UserID, req.ID)
		response := PostResponse{Posts: nil, Success: 1, Message: "Post Deleted!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}
