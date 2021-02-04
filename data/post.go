package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	ID     string
	UserID string
	Images []string
	Date   string
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
	posts, err := findPosts(ctx, "postid", postID)

	return posts, err
}

func CreatePost(ctx context.Context, p Post) error {
	p.Date = time.Now().UTC().String()
	_, err := collection.InsertOne(ctx, p)

	return err
}

func DeletePost(ctx context.Context, postID string) error {
	_, err := collection.DeleteOne(ctx, bson.M{"postid": postID})

	return err
}
