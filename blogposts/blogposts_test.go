package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/sophiabrandt/learn-go-with-tests/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("returns blog posts from filesystem", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, terminal`
		)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := blogposts.Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}}

		assertPost(t, got, want)
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
