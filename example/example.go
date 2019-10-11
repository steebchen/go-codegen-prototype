package example

import (
	"context"
	"log"

	. "go-codegen/example/photon"
)

func example() {
	client := NewClient()

	var result UserQueryA
	err := client.Post.SelectParent.Name("UserQueryA").Fields(
		Post.ID.Select(),
		Post.Title.Select(),
		Post.Likes.Sum(),
		Post.Count(),
	).GroupBy(
		Post.ID.Group(),
	).Into(&result).Exec(context.Background())

	log.Printf("err: %s", err)
}

type UserQueryA struct {
	ID        string
	Title     string
	Content   string
	Likes     string
	PostCount string
}
