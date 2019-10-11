package main

import (
	"context"
	"log"

	. "go-codegen/example/photon"
)

func main() {
	client := NewClient()

	// sql equivalent:
	// SELECT sum(likes), count(id)
	// FROM Posts
	// GROUP BY title

	var result []Query1
	err := client.Post.Select.Name("Query1").Fields(
		Post.Likes.Sum(),
		Post.Count(),
	).GroupBy(
		Post.Title.Group(),
	).Into(&result).Exec(context.Background())

	var countResult []CountQuery
	err = client.Post.Select.Name("CountQuery").Fields(
		Post.Count(),
	).Into(&countResult).Exec(context.Background())

	if err != nil {
		panic(err)
	}

	for _, item := range result {
		log.Printf("item: %+v", item)
		log.Printf("title: %s", item.Title)
		log.Printf("likes: %d", item.Likes)
		log.Printf("posts: %d", item.PostCount)
	}

	log.Printf("count: %s", countResult[0].PostCount)
}
