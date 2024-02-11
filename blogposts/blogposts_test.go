package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/sophiabrandt/learn-go-with-tests/blogposts"
)

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
A
B
`
	secondBody = `Title: Post 2
Description: Description 2
Tags: rust, terminal
---
Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.
Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.
`
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("returns blog posts from filesystem", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		expected := []blogposts.Post{
			{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: "A\nB"},
			{Title: "Post 2", Description: "Description 2", Tags: []string{"rust", "terminal"}, Body: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.\nLorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(expected) {
			t.Errorf("got %d posts, want %d", len(posts), len(expected))
		}

		for i, post := range posts {
			assertPost(t, post, expected[i])
		}
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
