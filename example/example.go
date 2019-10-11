package example

import (
	"context"
	"log"

	. "go-codegen/example/photon"
)

func example() {
	client := NewClient()

	query, err := client.Post.SelectParent.Name("UserQueryA").Fields(
		Post.Title.Select(),
		Post.Likes.Sum(),
		Post.Count(),
	).GroupBy(
		Post.ID.Group(),
	).Exec(context.Background())

	if err != nil {
		panic(err)
	}

	log.Printf("query: %s", query)
	log.Printf("title: %s", query.Title)
	log.Printf("likes: %s", query.Likes)
	log.Printf("posts: %s", query.PostCount)
}
