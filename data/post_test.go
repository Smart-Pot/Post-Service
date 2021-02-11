package data

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Smart-Pot/pkg"
)

var testPostID string
var testUserID string

func TestMain(m *testing.M) {
	wd, _ := os.Getwd()
	pkg.ConfigOptions.BaseDir = filepath.Join(wd, "..", "config")
	err := pkg.Config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	DatabaseConnection()
	m.Run()
}

func TestCreatePost(t *testing.T) {
	testPost := Post{
		UserID: "test",
		Plant:  "test",
		Info:   "test",
		Images: []string{"test", "test"},
		EnvData: EnvData{
			Humidity:    "test",
			Temperature: "test",
			Light:       "test",
		},
	}
	testUserID = testPost.UserID

	err := CreatePost(nil, testPost)

	if err != nil {
		t.Error(err)
	}
}

func TestGetUsersPosts(t *testing.T) {
	results, err := GetUsersPosts(nil, testUserID)

	if err != nil {
		t.Error(err)
	}
	if len(results) <= 0 {
		t.Errorf("Post not found! UserID = %s", testUserID)
		t.FailNow()
	}
	result := results[0]

	if result.UserID != testUserID {
		t.Errorf("Expected userID = %s, found = %s", result.UserID, testUserID)
	}
	testPostID = result.ID
}

func TestGetPost(t *testing.T) {
	results, err := GetPost(nil, testPostID)

	if err != nil {
		t.Error(nil)
	}
	if len(results) <= 0 {
		t.Errorf("Post not found! postID = %s", testPostID)
		t.FailNow()
	}

	result := results[0]
	if result.ID != testPostID {
		t.Errorf("Expected postID = %s, found = %s", result.ID, testPostID)
	}
}

func TestVote(t *testing.T) {
	err := Vote(nil, testUserID, testPostID)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateLikes(t *testing.T) {
	var likes []string
	likes = updateLikes("test", likes)
	if len(likes) <= 0 {
		t.Error("Fail adding to array!")
		t.FailNow()
	}
	if likes[0] != "test" {
		t.Error("Wrong string added to array!")
		t.FailNow()
	}

	likes = updateLikes("test", likes)
	if len(likes) != 0 {
		t.Errorf("Dislike failed! likes = %s", likes)
	}
}

func TestDeletePost(t *testing.T) {
	err := DeletePost(nil, testPostID)

	if err != nil {
		t.Error(err)
	}
}
