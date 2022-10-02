package main

import (
	"fmt"

	"github.com/hadihammurabi/belajar-go-goroutine/blocking"
	"github.com/hadihammurabi/belajar-go-goroutine/utils"
)

func main() {
	go func() {
		ids, err := utils.GetPostIds()
		if err != nil {
			panic(err)
		}
		fmt.Println(ids)
	}()

	go func() {
		post, err := utils.GetPostDetail(1)
		if err != nil {
			panic(err)
		}
		fmt.Println(post)
	}()

	blocking.Forever()
}
