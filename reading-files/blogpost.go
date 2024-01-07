package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator = "Title: "
	descSeparator  = "Description: "
	tagsSeparator  = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return createPost(postFile)
}

func createPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readLine()[len(titleSeparator):]
	description := readLine()[len(descSeparator):]
	tagsLine := readLine()[len(tagsSeparator):]
	tags := strings.Split(tagsLine, ", ")
	body := readBody(scanner)

	return Post{Title: title, Description: description, Tags: tags, Body: body}, nil
}

func readBody(scanner *bufio.Scanner) string {
	// --- before body
	scanner.Scan()

	bodyBuffer := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&bodyBuffer, scanner.Text())
	}
	return strings.TrimSuffix(bodyBuffer.String(), "\n")
}
