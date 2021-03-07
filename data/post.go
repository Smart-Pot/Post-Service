package data

import (
	"context"
	"errors"
	"time"

	"github.com/Smart-Pot/pkg/db"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID      string   `json:"id"`
	UserID  string   `json:"userId" validate:"required"`
	Plant   string   `json:"plant" validate:"required"`
	Info    string   `json:"info" validate:"required"`
	EnvData EnvData  `json:"envData" validate:"required"`
	Images  []string `json:"images" validate:"required"`
	Like    []string `json:"like"`
	Deleted bool     `json:"deleted"`
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

func findPosts(ctx context.Context, filter interface{}, opts *options.FindOptions) ([]*Post, error) {
	var results []*Post
	cur, err := db.Collection().Find(ctx, filter, opts)
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

func GetUsersPosts(ctx context.Context, userID string, pageNumber, pageSize int) ([]*Post, error) {
	skip := int64((pageNumber - 1) * pageSize)
	limit := int64(pageSize)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	filter := bson.M{
		"userid":  userID,
		"deleted": false,
	}
	posts, err := findPosts(ctx, filter, &opts)

	return posts, err
}

func GetPost(ctx context.Context, postID string) ([]*Post, error) {
	filter := bson.M{
		"postid":  postID,
		"deleted": false,
	}
	posts, err := findPosts(ctx, filter, nil)

	if len(posts) <= 0 {
		return nil, errors.New("post not found")
	}

	return posts, err
}

func CreatePost(ctx context.Context, p Post) error {
	p.Date = time.Now().UTC().String()
	p.ID = generateID()
	p.Like = []string{}
	p.Deleted = false
	_, err := db.Collection().InsertOne(ctx, p)

	return err
}

func Vote(ctx context.Context, userID string, postID string) error {
	res := db.Collection().FindOne(ctx, bson.M{"id": postID})
	var p Post
	if err := res.Decode(&p); err != nil {
		return err
	}
	p.Like = updateLikes(userID, p.Like)
	filter := bson.M{"id": postID}
	pushToArray := bson.M{"$set": bson.M{"like": p.Like}}
	result, err := db.Collection().UpdateOne(ctx, filter, pushToArray)
	if result.ModifiedCount <= 0 {
		return errors.New("vote failed")
	}
	return err
}

func updateLikes(userID string, likes []string) []string {
	for i, v := range likes {
		if v == userID {
			return append(likes[:i], likes[i+1:]...)
		}
	}
	return append(likes, userID)
}

func DeletePost(ctx context.Context, postID string) error {
	filter := bson.M{"id": postID}

	updatePost := bson.M{"$set": bson.M{"deleted": true}}

	res, err := db.Collection().UpdateOne(ctx, filter, updatePost)
	if err != nil {
		return err
	}

	if res.ModifiedCount <= 0 {
		return errors.New("post not found")
	}

	return nil
}

func DeletePosts(ctx context.Context, userID string) error {
	filter := bson.M{"userid": userID}
	updatePost := bson.M{"$set": bson.M{"deleted": true}}
	_, err := db.Collection().UpdateMany(ctx, filter, updatePost)

	if err != nil {
		return err
	}

	return nil
}

func generateID() string {
	return uuid.NewString()
}
