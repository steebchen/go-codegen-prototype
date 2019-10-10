package example

import (
	"context"
	"log"

	. "go-codegen/example/photon"
)

func example() {
	ctx := context.Background()
	client := NewClient()

	var result interface{}
	err := client.Post.Select.Name("UserQueryA").GroupBy(
		Post.Title.Group(),
	).Fields(
		Post.Count(),
		Post.Likes.Sum(),
	).Into(&result).Exec(ctx)

	log.Printf("err: %s", err)
}
