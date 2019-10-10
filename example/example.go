package example

import (
	"context"
	"log"

	. "go-codegen/example/photon"
)

func example() {
	client := NewClient()

	var result interface{}
	err := client.Post.Select.Name("UserQueryA").GroupBy(
		Post.Title.Group(),
	).Fields(
		Post.Count(),
		Post.Likes.Sum(),
	).Into(&result).Exec(context.Background())

	log.Printf("err: %s", err)
}
