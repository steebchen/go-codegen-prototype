package main

import (
	"context"
	"log"

	. "go-codegen/example/photon"
)

func main() {
	client := NewClient()

	query, err := client.Post.Select.Name("UserQueryA").Fields(
		Post.Likes.Sum(),
		Post.Count(),
	).GroupBy(
		Post.Title.Group(),
	).Exec(context.Background())

	if err != nil {
		panic(err)
	}

	for _, item := range query {
		log.Printf("item: %+v", item)
		log.Printf("title: %s", item.Title)
		log.Printf("likes: %d", item.Likes)
		log.Printf("posts: %d", item.PostCount)
	}
}
