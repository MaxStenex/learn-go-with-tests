package blogposts

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(`Title: Post 1
Description: Desc 1
Tags: one, two
---
Hello
World
!`)},
		"hello-world2.md": {Data: []byte(`Title: Post 2
Description: Desc 2,
Tags: three
---`)},
	}

	posts, _ := NewPostsFromFS(fs)

	got := posts[0]
	want := Post{Title: "Post 1", Description: "Desc 1", Tags: []string{"one", "two"}, Body: `Hello
World
!`}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, \n want %+v", got, want)
	}
}
