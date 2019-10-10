package example

import (
	"context"
	"log"

	. "go-codegen/example/photon"
)

func example() {
	client := NewClient()

	var result interface{}
	err := client.Post.SelectParent.Name("UserQueryA").Fields(
		Post.ID.Group(),
		Post.Title.Select(),
	).GroupBy(
		Post.Content.Select(),
		Post.Likes.Sum(),
	).Into(&result).Exec(context.Background())

	log.Printf("err: %s", err)
}
