package main

import (
	"context"
	"log"

	. "go-codegen/example/photon"
)

func main() {
	client := NewClient()
	ctx := context.Background()

	// sql equivalent:
	// SELECT sum(likes), count(id)
	// FROM Posts
	// GROUP BY title

	var result Query1
	err := client.Post.Select.Name("Query1").Fields(
		Post.Likes.Sum(),
		Post.Count(),
	).GroupBy(
		Post.Title.Group(),
	).Into(&result).Exec(ctx)

	if err != nil {
		panic(err)
	}

	for _, item := range result {
		log.Printf("item: %+v", item)
		log.Printf("likes: %d", item.Likes)
		log.Printf("posts: %d", item.PostCount)
	}
}
