package data

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	ID      string   `json:"id"`
	UserID  string   `json:"userId" validate:"required"`
	EnvData EnvData  `json:"envData" validate:"required"`
	Images  []string `json:"images" validate:"required"`
	Date    string   `json:"date"`
}

type EnvData struct {
	Humidity    string `json:"humidity" validate:"required"`
	Temperature string `json:"temperature" validate:"required"`
	Light       string `json:"light" validate:"required"`
}

func (p *Post) Validate() error {
	v := validator.New()
	return v.Struct(p)
}

func findPosts(ctx context.Context, key, value string) ([]*Post, error) {
	var results []*Post
	cur, err := collection.Find(ctx, bson.D{{key, value}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var pst Post
		err := cur.Decode(&pst)
		if err != nil {
			return nil, err
		}

		results = append(results, &pst)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	return results, err
}

func GetUsersPosts(ctx context.Context, userID string) ([]*Post, error) {
	posts, err := findPosts(ctx, "userid", userID)

	return posts, err
}

func GetPost(ctx context.Context, postID string) ([]*Post, error) {
	posts, err := findPosts(ctx, "id", postID)

	if len(posts) <= 0 {
		return nil, errors.New("post not found")
	}

	return posts, err
}

func CreatePost(ctx context.Context, p Post) error {
	p.Date = time.Now().UTC().String()
	p.ID = generateID()
	_, err := collection.InsertOne(ctx, p)

	return err
}

func DeletePost(ctx context.Context, postID string) error {
	r, err := collection.DeleteOne(ctx, bson.M{"id": postID})
	if r.DeletedCount <= 0 {
		return errors.New("post not found")
	}
	return err
}

func generateID() string {
	return uuid.NewString()
}
