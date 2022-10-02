package main

import (
	"fmt"

	"github.com/hadihammurabi/belajar-go-goroutine/utils"
)

func main() {
	postIdsChan := make(chan []int)
	go func() {
		postIds, err := utils.GetPostIds()
		if err != nil {
			panic(err)
		}

		postIdsChan <- postIds
	}()

	postsChan := make(chan utils.Post)
	go func() {
		postIds := <-postIdsChan
		for _, id := range postIds {
			go func(id int) {
				post, err := utils.GetPostDetail(id)

				if err == nil {
					postsChan <- post
				}
			}(id)
		}
	}()

	for post := range postsChan {
		fmt.Println(post)
	}

	// blocking.Forever()
}
