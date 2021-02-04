package service

import (
	"context"
	"postservice/data"

	"github.com/go-kit/kit/log"
)

type service struct {
	logger log.Logger
}

type Service interface {
	GetSingle(ctx context.Context, postID string) ([]*data.Post, error)
	GetMultiple(ctx context.Context, userID string) ([]*data.Post, error)
	Create(ctx context.Context, newPost data.Post) error
	Delete(ctx context.Context, postID string) error
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) GetSingle(ctx context.Context, postID string) ([]*data.Post, error) {
	return data.GetPost(ctx, postID)
}
func (s service) GetMultiple(ctx context.Context, userID string) ([]*data.Post, error) {
	return data.GetUsersPosts(ctx, userID)
}

func (s service) Create(ctx context.Context, newPost data.Post) error {
	return data.CreatePost(ctx, newPost)
}

func (s service) Delete(ctx context.Context, postID string) error {
	return data.DeletePost(ctx, postID)
}
