package utils

import (
	"fmt"

	"github.com/goreq/goreq"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func GetPostIds() ([]int, error) {
	res, err := goreq.Get("https://jsonplaceholder.typicode.com/posts", nil)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	posts := make([]Post, 0)
	err = res.Json(&posts)
	if err != nil {
		return nil, err
	}

	postIds := make([]int, 0)
	for _, post := range posts {
		postIds = append(postIds, post.ID)
	}

	return postIds, nil
}

func GetPostDetail(id int) (Post, error) {
	res, err := goreq.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id), nil)

	if err != nil {
		return Post{}, err
	}

	defer res.Body.Close()

	post := Post{}
	err = res.Json(&post)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}
