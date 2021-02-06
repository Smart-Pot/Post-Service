package service

import (
	"context"
	"errors"
	"postservice/data"
	"time"

	"github.com/Smart-Pot/pkg/adapter/amqp"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	logger   log.Logger
	producer amqp.Producer
}

type Service interface {
	GetSingle(ctx context.Context, postID string) ([]*data.Post, error)
	GetMultiple(ctx context.Context, userID string) ([]*data.Post, error)
	Create(ctx context.Context, userID string, newPost data.Post) error
	Delete(ctx context.Context, userID, postID string) error
	Vote(ctx context.Context, userID, postID string) error
}

func NewService(logger log.Logger) Service {
	p, err := amqp.MakeProducer("DeletePostComments")
	if err != nil {
		panic(err)
	}
	return &service{
		logger:   logger,
		producer: p,
	}
}

func (s service) Vote(ctx context.Context, userID, postID string) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Vote",
			"param:userID", userID,
			"param:commentID", postID,
			"took", time.Since(beginTime))
	}(time.Now())
	return data.Vote(ctx, userID, postID)
}

func (s service) GetSingle(ctx context.Context, postID string) (result []*data.Post, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "GetSingle",
			"param:postID", postID,
			"result", result,
			"took", time.Since(beginTime),
			"error", err,
		)
	}(time.Now())
	result, err = data.GetPost(ctx, postID)
	return result, err
}
func (s service) GetMultiple(ctx context.Context, userID string) (result []*data.Post, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "GetMultiple",
			"param:userID", userID,
			"result", result,
			"took", time.Since(beginTime),
			"error", err,
		)
	}(time.Now())
	result, err = data.GetUsersPosts(ctx, userID)
	return result, err
}

func (s service) Create(ctx context.Context, userID string, newPost data.Post) (err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Create",
			"param:newPost", newPost,
			"param:userID", userID,
			"took", time.Since(beginTime),
			"error", err,
		)
	}(time.Now())
	if err := newPost.Validate(); err != nil {
		return err
	}
	if newPost.UserID != userID {
		return errors.New("User can not create comments for other users")
	}
	return data.CreatePost(ctx, newPost)
}

func (s service) Delete(ctx context.Context, userID, postID string) (err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Delete",
			"param:postID", postID,
			"param:userID", userID,
			"took", time.Since(beginTime),
			"error", err,
		)
	}(time.Now())
	posts, err := data.GetPost(ctx, postID)
	if err != nil {
		return err
	}
	post := posts[0]
	if userID != post.UserID {
		return errors.New("User can not delete comments of other users")
	}
	err = data.DeletePost(ctx, postID)
	if err != nil {
		return err
	}
	s.producer.Produce([]byte(postID))
	return err
}
